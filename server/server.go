package main

import (
	"context"
	"time"

	"github.com/brevis-network/uniswap-rebate/binding"
	"github.com/brevis-network/uniswap-rebate/circuit"
	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/brevis-network/uniswap-rebate/onchain"
	"github.com/brevis-network/uniswap-rebate/webapi"
	"github.com/celer-network/goutils/log"
)

type Server struct {
	db *dal.DAL
}

func (s *Server) NewProof(_ context.Context, _ *webapi.NewProofReq) (*webapi.NewProofResp, error) {
	return &webapi.NewProofResp{
		Errmsg: "manual proof request is disabled; backend generates proofs on schedule",
	}, nil
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
	ret.Calldata = calldata
	return ret, nil
}

func (s *Server) RunScheduledProver(ctx context.Context, intv time.Duration) {
	ticker := time.NewTicker(intv)
	defer ticker.Stop()

	s.RunScheduledOnce(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.RunScheduledOnce(ctx)
		}
	}
}

func (s *Server) RunScheduledOnce(ctx context.Context) {
	rows, err := s.db.Claimers(ctx)
	if err != nil {
		log.Error("db Claimers err:", err)
		return
	}
	for _, row := range rows {
		s.runOneRouter(ctx, row)
	}
}

func (s *Server) runOneRouter(ctx context.Context, row dal.Claimer) {
	onec, ok := chainMap[row.Chid]
	if !ok {
		log.Warnf("skip unsupported source chain %d for router %s", row.Chid, row.Router)
		return
	}
	safeBlk, err := onec.LatestSafeBlock(ctx)
	if err != nil {
		log.Errorf("LatestSafeBlock chain %d err: %v", row.Chid, err)
		return
	}
	// First time seeing this router: initialize cursor to current safe block.
	// This avoids scanning full chain history.
	if row.FetchBlk == 0 {
		err = s.db.ClaimerSetFetchBlk(ctx, dal.ClaimerSetFetchBlkParams{
			FetchBlk: safeBlk,
			Chid:     row.Chid,
			Router:   row.Router,
		})
		if err != nil {
			log.Errorf("init ClaimerSetFetchBlk chain %d router %s err: %v", row.Chid, row.Router, err)
		}
		return
	}
	fromBlk := row.FetchBlk + 1
	if fromBlk > safeBlk {
		return
	}

	router := onchain.Hex2addr(row.Router)
	// todo: proper end block
	swaps, err := onec.FetchRouterSwaps(router, fromBlk, fromBlk+100)
	if err != nil {
		log.Errorf("FetchRouterSwaps chain %d router %s [%d,%d] err: %v", row.Chid, row.Router, fromBlk, safeBlk, err)
		return
	}
	swaps = filterOversizedBlocks(swaps)
	if len(swaps) > 0 {
		reqid := time.Now().UnixNano()
		info := binding.ProofInfo{
			ReqId:      reqid,
			ChainId:    row.Chid,
			PoolMgr:    onec.PoolMgr,
			GasPerSwap: onec.GasPerSwap,
			GasPerTx:   onec.GasPerTx,
			Logs:       swaps,
		}
		err = s.db.ReqAdd(ctx, dal.ReqAddParams{
			ID:        reqid,
			Router:    row.Router,
			UsrReq:    nil,
			ProofInfo: info,
		})
		if err != nil {
			log.Errorf("ReqAdd reqid %d err: %v", reqid, err)
			return
		}
		log.Infof("start proving reqid %d chain %d router %s swaps %d", reqid, row.Chid, row.Router, len(swaps))
		go prMgr.Run(&info)
	}
	/*
		// processed this range, move cursor forward regardless of swaps found.
		err = s.db.ClaimerSetFetchBlk(ctx, dal.ClaimerSetFetchBlkParams{
			FetchBlk: safeBlk,
			Chid:     row.Chid,
			Router:   row.Router,
		})
		if err != nil {
			log.Errorf("ClaimerSetFetchBlk chain %d router %s err: %v", row.Chid, row.Router, err)
		}
	*/
}

func filterOversizedBlocks(swaps []binding.OneLog) []binding.OneLog {
	if len(swaps) == 0 {
		return swaps
	}
	blkNums, blk2swaps := binding.SwapsByBlock(swaps)
	ret := make([]binding.OneLog, 0, len(swaps))
	for _, blknum := range blkNums {
		group := blk2swaps[blknum]
		if len(group.PoolIds) > circuit.MaxPoolNum {
			log.Warnf("skip block %d: pools %d > max %d", blknum, len(group.PoolIds), circuit.MaxPoolNum)
			continue
		}
		if len(group.Logs) > circuit.MaxSwapNum {
			log.Warnf("skip block %d: swaps %d > max %d", blknum, len(group.Logs), circuit.MaxSwapNum)
			continue
		}
		ret = append(ret, group.Logs...)
	}
	return ret
}
