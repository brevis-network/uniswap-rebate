package proofmgr

import (
	"context"

	"github.com/brevis-network/brevis-sdk/sdk/proto/commonproto"
	"github.com/brevis-network/brevis-sdk/sdk/proto/gwproto"
	"github.com/brevis-network/brevis-sdk/sdk/proto/sdkproto"
	"github.com/brevis-network/uniswap-rebate/binding"
	"github.com/brevis-network/uniswap-rebate/circuit"
	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/celer-network/goutils/log"
	"google.golang.org/grpc"
)

const (
	BrvGwApiKey = "UniGas" // UniGasAgg
)

type ProofMgr struct {
	gwclient   gwproto.GatewayClient
	db         *dal.DAL
	dstChId    uint64 // 130 Unichain mainnet
	appProvers []string
}

func NewProofMgr(gwrpc string, appProvers []string, dstChId uint64, db *dal.DAL) *ProofMgr {
	conn, _ := grpc.NewClient(gwrpc)
	return &ProofMgr{
		appProvers: appProvers,
		dstChId:    dstChId,
		db:         db,
		gwclient:   gwproto.NewGatewayClient(conn),
	}
}

// blocking until final proofs are saved
// 1. ProofInfo -> []sdkproto.ProveRequest. one ProveRequest corresponds to one app proof and one gw Query
// 2. iter reqs. send to app prover and save appCircuitInfo
// 3. build Queries from reqs and received appCircuitInfo, call gw.SendBatchQueriesAsync
// 4. for each req, polling app proof and gw query to see if both are ready
// 5. if yes, call gw.SubmitAppProof
func (m *ProofMgr) Run(info *binding.ProofInfo) {
	sdkreqs := m.BuildProveReqs(info)
	var appInfos []*commonproto.AppCircuitInfo
	// ===== app prover
	for idx, req := range sdkreqs {
		appInfo, err := m.DoAppProveAsync(info.ReqId, idx, m.appProvers[idx%len(m.appProvers)], req)
		if err != nil {
			log.Errorln(info.ReqId, idx, "DoAppProveAsync err:", err)
			return
		}
		appInfos = append(appInfos, appInfo)
	}
	// ===== gw
	gwreq := &gwproto.SendBatchQueriesRequest{
		ChainId:       info.ChainId,
		TargetChainId: m.dstChId,
		Option:        gwproto.QueryOption_ZK_MODE,
		ApiKey:        BrvGwApiKey,
	}
	gwreq.Queries = buildGwQueries(sdkreqs, appInfos)
	asyncResp, err := m.gwclient.SendBatchQueriesAsync(context.Background(), gwreq)
	if err != nil {
		log.Errorln(info.ReqId, "SendBatchQueriesAsync err:", err)
		return
	}
	if asyncResp.Err != nil {
		log.Errorln(info.ReqId, "SendBatchQueriesAsync resp.Err:", asyncResp.Err)
		return
	}
	// save to db
	for idx, gwreqid := range asyncResp.RequestIds {
		m.db.ProofSetGwInfo(context.Background(), dal.ProofSetGwInfoParams{
			GatewayBatchID:   asyncResp.BatchId,
			GatewayRequestID: gwreqid,
			GatewayNonce:     asyncResp.Nonce,
			Reqid:            info.ReqId, // our own reqid
			Idx:              idx,        // nth proof/query for this Reqid
		})
	}
	// block parallel polling app prover status and gw
}

// proofinfo -> list of ProveRequest
func (m *ProofMgr) BuildProveReqs(info *binding.ProofInfo) (ret []*sdkproto.ProveRequest) {
	claimLog := info.Logs[0]
	swaps := info.Logs[1:]
	// split swaps into groups without exceeding circuit max. each group needs a separate app proof and gw Query
	for _, swapGroup := range binding.SplitIntoGroups(swaps, circuit.MaxSwapNum, circuit.MaxPoolNum) {
		req := &sdkproto.ProveRequest{
			SrcChainId: info.ChainId,
		}
		req.Receipts = append(req.Receipts, evToIndexedReceipt(claimLog, 0))
		for i, ev := range swapGroup.Logs {
			req.Receipts = append(req.Receipts, evToIndexedReceipt(ev, i+1))
		}
		appCirc := binding.NewCircuit(info, swapGroup.Logs, m.db.GetPoolKeys(info.ChainId, swapGroup.PoolIds))
		req.CustomInput, _ = buildCustomInput(appCirc)
		ret = append(ret, req)
	}
	return ret
}

// only topic 1 and topic 2 as fields. both our Claimer and uniswap swap are the same logic
func evToIndexedReceipt(ev binding.OneLog, index int) *sdkproto.IndexedReceipt {
	return &sdkproto.IndexedReceipt{
		Index: uint32(index),
		Data: &sdkproto.ReceiptData{
			BlockNum: ev.BlockNumber,
			TxHash:   ev.TxHash.Hex(),
			Fields: []*sdkproto.Field{
				{
					Contract:   ev.Address.Hex(),
					LogPos:     uint32(ev.Index - ev.LogIdxOffset),
					EventId:    ev.Topics[0].Hex(),
					Value:      ev.Topics[1].Hex(),
					IsTopic:    true,
					FieldIndex: 1,
				},
				{
					Contract:   ev.Address.Hex(),
					LogPos:     uint32(ev.Index - ev.LogIdxOffset),
					EventId:    ev.Topics[0].Hex(),
					Value:      ev.Topics[2].Hex(),
					IsTopic:    true,
					FieldIndex: 2,
				},
			},
		},
	}
}

// call appprover.ProveAsync, save app circuit info and proof id to db
func (m *ProofMgr) DoAppProveAsync(reqid int64, idx int, appProver string, proveReq *sdkproto.ProveRequest) (*commonproto.AppCircuitInfo, error) {
	client, _ := getProverClient(appProver)
	proverResp, err := client.ProveAsync(context.Background(), proveReq)
	if err != nil {
		log.Errorf("AppProveAsync fail, err: %v", err)
		return nil, err
	}
	if proverResp.Err != nil {
		log.Errorf("AppProveAsync proverResp err: %v", proverResp.Err)
		return nil, err
	}
	err = m.db.ProofAdd(context.Background(), dal.ProofAddParams{
		Reqid:          reqid,
		Idx:            idx,
		AppProver:      appProver,
		AppProofID:     proverResp.ProofId,
		AppCircuitInfo: proverResp.CircuitInfo,
	})
	return proverResp.CircuitInfo, err
}

// BlockPolling for one app prover and gw query, will return if both are true
func (m *ProofMgr) BlockPolling(appProofId, gwReqId string, gwNonce uint64) {

}
