/*
Copyright © 2024 Brevis Network
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/brevis-network/brevis-sdk/sdk/proto/gwproto"
	"github.com/brevis-network/uniswap-rebate/circuit"
	"github.com/brevis-network/uniswap-rebate/onchain"
	"github.com/celer-network/goutils/log"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/constraint"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/cobra"
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

// proveCmd represents the prove command
var proveCmd = &cobra.Command{
	Use:   "prove",
	Short: "accept prove req, send to Brevis",

	Run: func(cmd *cobra.Command, args []string) {
		// takes minutes to load
		var err error
		compiledCircuit, pk, vk, vkHash, err = sdk.ReadSetupFrom(outDir, circuit.MaxReceipts, circuit.MaxReceipts, circuit.MaxReceipts*2)
		chkErr(err, "ReadSetupFrom "+outDir)
		vkhash, err := sdk.ComputeVkHash(vk)
		chkErr(err, "ComputeVkHash")
		vkHashStr = fmt.Sprintf("%s", vkhash) // common.Hash.Format adds 0x prefix
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

func ProveReqs(reqs []*onchain.OneProveReq) {
	for _, r := range reqs {
		DoOne(r)
		// query, err := DoOne(r)

	}
}

func DoOne(r *onchain.OneProveReq) (*gwproto.Query, error) {
	log.Infoln("req", r.ReqId, "pools:", r.PoolIds)

	app, _ := sdk.NewBrevisApp(r.ChainId, brvGw)
	// var receipts []*sdk.ReceiptData
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
		// app.AddReceipt()
	}
	c := r.NewCircuit()
	_, err := app.BuildCircuitInput(c)
	if err != nil {
		return nil, fmt.Errorf("BuildCircuitInput %v", err)
	}
	return nil, nil
}

func init() {
	rootCmd.AddCommand(proveCmd)
	proveCmd.PersistentFlags().StringVar(&outDir, "outDir", "$HOME/circuitOut/unigasrebate", "folder for circuit in/output")
	proveCmd.PersistentFlags().StringVar(&brvGw, "brvgw", "", "brevis gateway grpc endpoint")
	proveCmd.PersistentFlags().IntVar(&lport, "port", 8889, "listen port for prove request")
}

func chkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
