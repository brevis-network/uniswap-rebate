/*
Copyright © 2024 Brevis Network
*/
package cmd

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"path/filepath"
	"time"

	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/brevis-network/brevis-sdk/sdk/proto/commonproto"
	"github.com/brevis-network/brevis-sdk/sdk/proto/gwproto"
	"github.com/brevis-network/uniswap-rebate/circuit"
	"github.com/brevis-network/uniswap-rebate/onchain"
	"github.com/celer-network/goutils/log"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/constraint"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// flags
var (
	brvGw string
	// accept req on this port
	lport int
)
var (
	compiledCircuit constraint.ConstraintSystem
	pk              plonk.ProvingKey
	vk              plonk.VerifyingKey
	vkHashStr       string
	vkHash          []byte
)

const (
	BrvApiKey = "123456" // need another key?
)

// proveCmd represents the prove command
var proveCmd = &cobra.Command{
	Use:   "prove",
	Short: "accept prove req, send to Brevis",

	Run: func(cmd *cobra.Command, args []string) {
		appCircuit := circuit.DefaultCircuit()

		// takes minutes to load
		var err error
		compiledCircuit, pk, vk, vkHash, err = sdk.ReadSetupFrom(appCircuit, outDir)
		chkErr(err, "ReadSetupFrom "+outDir)
		// var buf bytes.Buffer
		// vk.WriteTo(&buf)
		// fmt.Printf("==== vk ====:\n%x", buf.Bytes())
		vkHashStr = fmt.Sprintf("0x%x", vkHash)
		log.Infoln("vkHashStr:", vkHashStr)
		// setup server
		router := httprouter.New()
		router.POST("/prove", HandleProve)
		log.Infoln("listen on port:", lport)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", lport), router))
	},
}

// handle StartProve req
func HandleProve(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req []*onchain.OneProveReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request: "+err.Error(), http.StatusBadRequest)
		return
	}
	go ProveReqs(req)
	fmt.Fprint(w, "OK\n")
}

// generate app proof, send to brevis gateway
func ProveReqs(reqs []*onchain.OneProveReq) {
	brvReq := &gwproto.SendBatchQueriesRequest{
		ChainId:       reqs[0].ChainId,
		TargetChainId: reqs[0].ChainId,
		Option:        gwproto.QueryOption_ZK_MODE, // 0 anyway so could omit
		ApiKey:        BrvApiKey,
	}
	for i, r := range reqs {
		q, err := DoOneReq(r, i)
		if err != nil {
			log.Errorln("batch", i, "err:", err)
			continue
		}
		brvReq.Queries = append(brvReq.Queries, q)
	}

	// submit
	conn, err := grpc.Dial(brvGw, grpc.WithBlock(), grpc.WithTimeout(5*time.Second), grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	if err != nil {
		log.Errorln("dial", brvGw, "err:", err)
		return
	}
	client := gwproto.NewGatewayClient(conn)
	log.Infoln("send", len(brvReq.Queries), "batches to brevis")
	resp, err := client.SendBatchQueries(context.Background(), brvReq)
	if err != nil {
		log.Errorln("SendBatchQueries err:", err)
		return
	}
	log.Info(resp)
}

func DoOneReq(r *onchain.OneProveReq, batchIdx int) (*gwproto.Query, error) {
	log.Infoln("req", r.ReqId, "pools:", r.PoolIds)

	app, _ := sdk.NewBrevisApp(r.ChainId, brvGw)
	var receipts []*sdk.ReceiptData
	var lastBlockNum uint64
	for i, onelog := range r.Logs {
		blkNum := onelog.Swap.BlockNumber
		block := r.Blks[blkNum]
		if blkNum != lastBlockNum {
			app.AddStorage(sdk.StorageData{
				BlockNum:     new(big.Int).SetUint64(blkNum),
				BlockBaseFee: new(big.Int).SetUint64(block.BaseFee),
				Address:      onchain.Hex2addr(r.Oracle),
				Slot:         onchain.ZeroHash,
				Value:        block.SlotValue,
			}, i) // special mode, sdk will fill dummy in between
		}
		rd := OneLog2SdkReceipt(onelog, block.BaseFee)
		receipts = append(receipts, &rd)
		app.AddReceipt(rd)
	}
	appCircuit := r.NewCircuit()
	circuitInput, err := app.BuildCircuitInput(appCircuit)
	if err != nil {
		return nil, fmt.Errorf("BuildCircuitInput %v", err)
	}
	witness, publicWitness, err := sdk.NewFullWitness(appCircuit, circuitInput)
	if err != nil {
		return nil, fmt.Errorf("NewFullWitness %v", err)
	}
	proof, err := sdk.Prove(compiledCircuit, pk, witness)
	if err != nil {
		return nil, fmt.Errorf("Prove %v", err)
	}
	err = sdk.WriteTo(proof, filepath.Join(outDir, fmt.Sprintf("%d-batch_%d.proof", r.ReqId, batchIdx)))
	if err != nil {
		return nil, fmt.Errorf("Write proof %v", err)
	}
	err = sdk.Verify(vk, publicWitness, proof)
	if err != nil {
		return nil, fmt.Errorf("Verify %v", err)
	}
	// bulid grpc Query
	var b bytes.Buffer
	proof.WriteTo(&b)
	proofStr := hex.EncodeToString(b.Bytes())
	b.Reset()
	publicWitness.WriteTo(&b)
	witnessStr := hex.EncodeToString(b.Bytes())
	return &gwproto.Query{
		AppCircuitInfo:    buildAppCircuitInfo(circuitInput, witnessStr, vkHashStr, proofStr, ""),
		ReceiptInfos:      buildReceiptInfos(receipts),
		StorageQueryInfos: buildStorageInfos(r.Blks, r.Oracle, onchain.Hash2Hex(onchain.ZeroHash)),
	}, nil
}

func OneLog2SdkReceipt(swap onchain.OneLog, basefee, timestamp uint64) sdk.ReceiptData {
	ret := sdk.ReceiptData{
		BlockNum:       new(big.Int).SetUint64(swap.BlockNumber),
		BlockBaseFee:   new(big.Int).SetUint64(basefee),
		BlockTimestamp: timestamp,
		TxHash:         swap.TxHash,
		MptKeyPath:     TxIdx2MptPath(swap.TxIndex),
	}
	ret.Fields[0] = sdk.LogFieldData{
		Contract:   swap.Address,
		LogPos:     swap.Index - swap.LogIdxOffset,
		EventID:    swap.Topics[0],
		IsTopic:    true,
		FieldIndex: 1,
		Value:      swap.Topics[1],
	}
	ret.Fields[1] = sdk.LogFieldData{
		Contract:   swap.Address,
		LogPos:     swap.Index - swap.LogIdxOffset,
		EventID:    swap.Topics[0],
		IsTopic:    true,
		FieldIndex: 2,
		Value:      swap.Topics[2],
	}
	return ret
}

func TxIdx2MptPath(txidx uint) *big.Int {
	var b []byte
	return new(big.Int).SetBytes(rlp.AppendUint64(b, uint64(txidx)))
}

func buildStorageInfos(m map[uint64]onchain.OneBlock, addr string, slot ...string) (infos []*gwproto.StorageQueryInfo) {
	for blknum := range m {
		infos = append(infos, &gwproto.StorageQueryInfo{
			Account:     addr,
			StorageKeys: slot,
			BlkNum:      blknum,
		})
	}
	return
}

func buildReceiptInfos(r []*sdk.ReceiptData) (infos []*gwproto.ReceiptInfo) {
	for _, d := range r {
		var logExtractInfo []*gwproto.LogExtractInfo
		for _, f := range d.Fields[:2] { // Fields is fixed length but we only need 2 logs, rest are empty and confuse server
			// we could also check for f.Contract isn't all 0
			logExtractInfo = append(logExtractInfo, &gwproto.LogExtractInfo{
				LogPos:         uint64(f.LogPos),
				ValueFromTopic: f.IsTopic,
				ValueIndex:     uint64(f.FieldIndex),
			})
		}
		infos = append(infos, &gwproto.ReceiptInfo{
			TransactionHash: d.TxHash.Hex(),
			LogExtractInfos: logExtractInfo,
		})
	}
	return
}

func buildAppCircuitInfo(in sdk.CircuitInput, witness, vk, proof, cbaddr string) *commonproto.AppCircuitInfoWithProof {
	inputCommitments := make([]string, len(in.InputCommitments))
	for i, value := range in.InputCommitments {
		inputCommitments[i] = fmt.Sprintf("0x%x", value)
	}

	toggles := make([]bool, len(in.Toggles()))
	for i, value := range in.Toggles() {
		toggles[i] = fmt.Sprintf("%x", value) == "1"
	}

	return &commonproto.AppCircuitInfoWithProof{
		OutputCommitment:  hex.EncodeToString(in.OutputCommitment.Hash().Bytes()),
		VkHash:            vk,
		InputCommitments:  inputCommitments,
		TogglesCommitment: fmt.Sprintf("0x%x", in.TogglesCommitment),
		Toggles:           toggles,
		Output:            hex.EncodeToString(in.GetAbiPackedOutput()),
		Proof:             proof,
		CallbackAddr:      cbaddr,
		// new fields required by plonky2
		InputCommitmentsRoot: fmt.Sprintf("0x%x", in.InputCommitmentsRoot),
		Witness:              witness,
		MaxReceipts:          circuit.MaxReceipts,
		MaxStorage:           circuit.MaxReceipts,
		MaxTx:                0,
		MaxNumDataPoints:     128, // hardcode for now
	}
}

func init() {
	rootCmd.AddCommand(proveCmd)
	proveCmd.PersistentFlags().StringVar(&outDir, "outDir", "$HOME/circuitOut/unigasrebate", "folder for circuit in/output")
	proveCmd.PersistentFlags().StringVar(&brvGw, "brvgw", "appsdkv3.brevis.network:443", "brevis gateway grpc endpoint")
	proveCmd.PersistentFlags().IntVar(&lport, "port", 9003, "listen port for prove request")
}

func chkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
