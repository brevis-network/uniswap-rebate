package proofmgr

import (
	"context"
	"time"

	"github.com/brevis-network/brevis-sdk/sdk/proto/commonproto"
	"github.com/brevis-network/brevis-sdk/sdk/proto/gwproto"
	"github.com/brevis-network/brevis-sdk/sdk/proto/sdkproto"
	"github.com/brevis-network/uniswap-rebate/binding"
	"github.com/brevis-network/uniswap-rebate/circuit"
	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/brevis-network/uniswap-rebate/webapi"
	"github.com/celer-network/goutils/log"
	"github.com/lthibault/jitterbug/v2"
	"golang.org/x/sync/errgroup"
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
// 4. for each req, polling app proof then call gw.SubmitAppProof, and polling gw till final proof
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
	// get all proof ids about this request and block polling app prover status and gw
	var errG errgroup.Group
	rows, _ := m.db.ProofGetIds(context.Background(), info.ReqId)
	for _, row := range rows {
		errG.Go(func() error {
			return m.DoOneProof(info.ReqId, row)
		})
	}
	if err := errG.Wait(); err != nil {
		log.Error("errG wait err:", err)
	}
}

// proofinfo -> list of ProveRequest
func (m *ProofMgr) BuildProveReqs(info *binding.ProofInfo) (ret []*sdkproto.ProveRequest) {
	// split swaps into groups without exceeding circuit max. each group needs a separate app proof and gw Query
	for _, swapGroup := range binding.SplitIntoGroups(info.Logs, circuit.MaxSwapNum, circuit.MaxPoolNum) {
		req := &sdkproto.ProveRequest{
			SrcChainId: info.ChainId,
		}
		for i, ev := range swapGroup.Logs {
			req.Receipts = append(req.Receipts, evToIndexedReceipt(ev, i))
		}
		appCirc := binding.NewCircuit(info, swapGroup.Logs, m.db.GetPoolKeys(info.ChainId, swapGroup.PoolIds))
		req.CustomInput, _ = buildCustomInput(appCirc)
		ret = append(ret, req)
	}
	return ret
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

// BLOCKING. polling app proof until ready, then call gw.SubmitAppProof, and polling gw till final proof
func (m *ProofMgr) DoOneProof(reqid int64, row dal.ProofGetIdsRow) error {
	client, _ := getProverClient(row.AppProver)
	var gwReq *gwproto.SubmitAppCircuitProofRequest
	t := jitterbug.New(
		30*time.Second,
		&jitterbug.Norm{Stdev: 10 * time.Second},
	)
	defer t.Stop() // avoid mem leak
	for range t.C {
		resp, _ := client.GetProof(context.Background(), &sdkproto.GetProofRequest{ProofId: row.AppProofID})
		if len(resp.Proof) == 0 {
			continue // try again next tick
		}
		// save app proof
		m.db.ProofSetAppProof(context.Background(), dal.ProofSetAppProofParams{
			AppProof:       resp.Proof,
			AppCircuitInfo: resp.CircuitInfo,
			AppProofID:     row.AppProofID,
		})
		// to be used in gw submit app proof
		gwReq = &gwproto.SubmitAppCircuitProofRequest{
			TargetChainId: m.dstChId,
			QueryKey: &gwproto.QueryKey{
				QueryHash: row.GatewayRequestID,
				Nonce:     row.GatewayNonce,
			},
			Proof:          resp.Proof,
			AppCircuitInfo: resp.CircuitInfo.ToInfoWithProof(resp.Proof, ""), // no callback addr
		}
		break // move to next step
	}
	// now try sending to gw. Note if gw side isn't ready, we keep retry
	for range t.C {
		resp, _ := m.gwclient.SubmitAppCircuitProof(context.Background(), gwReq)
		if resp.Err != nil {
			log.Errorln("submitAppProof err:", resp.Err)
		}
		if resp.Success {
			break // move to next step
		}
	}
	// now polling gw for final proof
	for range t.C {
		resp, _ := m.gwclient.GetQueryStatus(context.Background(), &gwproto.GetQueryStatusRequest{
			TargetChainId: m.dstChId,
			QueryKey: &gwproto.QueryKey{
				QueryHash: row.GatewayRequestID,
				Nonce:     row.GatewayNonce,
			},
		})
		if len(resp.Proof) > 0 {
			m.db.ProofSetGwResp(context.Background(), dal.ProofSetGwRespParams{
				GatewayQueryStatus: resp,
				GatewayRequestID:   row.GatewayRequestID,
				GatewayNonce:       row.GatewayNonce,
			})
			// TODO: handle agg proof of multiple queries
			m.db.ReqSetCalldata(context.Background(), dal.ReqSetCalldataParams{
				ID:       reqid,
				Calldata: webapi.GwRespToCalldata(resp),
			})
			break
		}
	}
	return nil
}
