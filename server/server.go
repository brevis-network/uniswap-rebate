package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/brevis-network/uniswap-rebate/binding"
	"github.com/brevis-network/uniswap-rebate/circuit"
	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/brevis-network/uniswap-rebate/onchain"
	"github.com/brevis-network/uniswap-rebate/webapi"
)

type Server struct {
	db *dal.DAL
}

func (s *Server) NewProof(ctx context.Context, req *webapi.NewProofReq) (ret *webapi.NewProofResp, err error) {
	reqid := time.Now().UnixMilli()
	router := onchain.Addr2hex(onchain.Hex2addr(req.Beneficiary))

	ret = new(webapi.NewProofResp)
	ret.Reqid = uint64(reqid)
	// parse and get tx receipt now even though it may take a long time for this api to return
	// so client knows error immediately instead of later polling
	onec, ok := chainMap[req.ChainId]
	if !ok {
		ret.Errmsg = fmt.Sprintf("unsupported chainid %d", req.ChainId)
		return
	}
	claimEv, err := s.db.ClaimerGet(context.Background(), dal.ClaimerGetParams{
		Chid:   req.ChainId,
		Router: router,
	})
	found, err := dal.ChkQueryRow(err)
	if !found {
		ret.Errmsg = fmt.Sprintf("unsupported router %s. Please contact Brevis team.", router)
		return
	}
	if err != nil {
		ret.Errmsg = fmt.Sprintf("db ClaimerGet err: %v", err)
		return
	}
	// fetch tx Receipts
	txList := strings.Split(req.TxnHashes, ",")
	log.Println("newproof", reqid, "chain:", req.ChainId, "router:", router, "tx num:", len(txList))
	receipts, err := onec.FetchTxReceipts(txList)
	if err != nil {
		ret.Errmsg = "get receipts err: " + err.Error()
		return
	}
	// check logs in receipts and prepare proof info
	swaps, err := onec.ProcessReceipts(receipts, onchain.Hex2addr(req.Beneficiary))
	if err != nil {
		ret.Errmsg = "ProcessReceipts err: " + err.Error()
		return
	}
	if len(swaps) == 0 {
		ret.Errmsg = "no eligible swaps found for " + router
		return
	}
	// Circuit supports up to MaxSwaps and MaxPoolNum
	// so we need to split into multiple requests if logs has more. Note we have to keep all swaps exactly ordered as onchain de-dup is by blknum
	// all swaps happen in same blk must be in one batch, so if MaxPool is 32 and within one block there are swaps with more than 32 pools
	// 33rd pool swaps will have no effect as circuit poolid check will return 0. So to avoid user confusion, we return error for now
	// TODO: handle exceed aggregation cap eg. 16 proofs
	blkNums, blk2swaps := binding.SwapsByBlock(swaps)
	for _, blknum := range blkNums { // blkNums is sorted ascending
		sameBlkSwaps := blk2swaps[blknum]
		// single block is too big, fail for now and guide requester to use another flow
		if len(sameBlkSwaps.PoolIds) > circuit.MaxPoolNum {
			ret.Errmsg = fmt.Sprintf("block %d includes %d pools, exceeds max %d", blknum, len(sameBlkSwaps.PoolIds), circuit.MaxPoolNum)
			return
		}
		if len(sameBlkSwaps.Logs) > circuit.MaxSwapNum {
			ret.Errmsg = fmt.Sprintf("block %d include %d swaps, exceeds max %d", blknum, len(sameBlkSwaps.Logs), circuit.MaxSwapNum)
			return
		}
	}

	info := binding.ProofInfo{
		ReqId:      reqid,
		ChainId:    req.ChainId,
		PoolMgr:    onec.PoolMgr,
		GasPerSwap: onec.GasPerSwap,
		GasPerTx:   onec.GasPerTx,
		Logs: []binding.OneLog{{
			Log:          &claimEv.Raw,
			LogIdxOffset: claimEv.Raw.Index,
		}},
	}
	info.Logs = append(info.Logs, swaps...)
	// save req to db
	err = s.db.ReqAdd(context.Background(), dal.ReqAddParams{
		ID:        reqid,
		Router:    router,
		UsrReq:    req,
		ProofInfo: info,
	})
	if err != nil {
		ret.Errmsg = "db ReqAdd err: " + err.Error()
		return
	}

	// enqueue

	// good to return
	return ret, nil
}

func (s *Server) GetProof(ctx context.Context, req *webapi.GetProofReq) (*webapi.GetProofResp, error) {
	ret := &webapi.GetProofResp{
		Reqid: req.Reqid,
	}
	calldata, err := s.db.ReqGetCalldata(context.Background(), int64(req.Reqid))
	if err != nil {
		return ret, err
	}
	// ret.Status
	ret.Calldata = calldata.ToWebCallData()
	return ret, nil
}
