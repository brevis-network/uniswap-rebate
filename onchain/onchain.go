package onchain

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/brevis-network/uniswap-rebate/binding"
	"github.com/brevis-network/uniswap-rebate/circuit"
	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	GetLogIntv   = time.Second * 60
	PoolInitEvId = "0xdd466e674ea557f56295e2d0218a125ea4b4f0f6f3307b95f85e6110838d6438"
	SwapEvId     = "0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f"
)

var (
	ZeroAddr common.Address
	ZeroHash common.Hash
)

type OneChain struct {
	*OneChainConfig
	ec  *ethclient.Client
	mon *mon2.Monitor
	db  *dal.DAL
}

type OneLog struct {
	*types.Log   // may also be Claimer event
	LogIdxOffset uint
	TxGasCap     uint32 // tx gas * 0.8. if 0, means next swap is from same tx
}

// only valid for swap, return topics[1]
func (o OneLog) PoolId() common.Hash {
	return o.Topics[1]
}

// return err if dial fail or chainid mismatch
func NewOneChain(cfg *OneChainConfig, dal *dal.DAL) (*OneChain, error) {
	ret := &OneChain{
		OneChainConfig: cfg,
		db:             dal,
	}
	var err error
	ret.ec, err = ethclient.Dial(cfg.Gateway)
	if err != nil {
		return nil, err
	}
	bgCtx := context.Background()
	chid, _ := ret.ec.ChainID(bgCtx)
	if chid == nil {
		return nil, fmt.Errorf("failed to retrieve rpc chid, cfg: %d", cfg.ChainID)
	}
	if chid.Uint64() != cfg.ChainID {
		return nil, fmt.Errorf("mismatch chainid cfg: %d, rpc: %d", cfg.ChainID, chid.Uint64())
	}

	ret.mon, err = mon2.NewMonitor(ret.ec, dal, mon2.PerChainCfg{
		BlkIntv:         time.Duration(cfg.BlkInterval) * time.Second,
		BlkDelay:        cfg.BlkDelay,
		MaxBlkDelta:     cfg.MaxBlkDelta,
		ForwardBlkDelay: cfg.ForwardBlkDelay,
	})
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *OneChain) Close() {
	c.mon.Close()
}

// txlist is hex string of tx hash, return non-nil err if any TransactionReceipt has err
// receipts are sorted by blocknum and index ascending
func (c *OneChain) FetchTxReceipts(txlist []string) ([]*types.Receipt, error) {
	var ret []*types.Receipt
	for _, tx := range txlist {
		r, err := c.ec.TransactionReceipt(context.Background(), Hex2hash(tx))
		if err != nil {
			return ret, err
		}
		ret = append(ret, r)
	}
	// sort by BlockNum and index ascending
	slices.SortFunc(ret, func(a, b *types.Receipt) int {
		blockNumCmp := a.BlockNumber.Cmp(b.BlockNumber)
		if blockNumCmp != 0 {
			return blockNumCmp
		}
		if a.TransactionIndex < b.TransactionIndex {
			return -1
		} else if a.TransactionIndex > b.TransactionIndex {
			return 1
		}
		return 0
	})
	return ret, nil
}

// go through receipt.Logs, check poolid is in db (eligible w/ non-zero hook addr) and sender matches
// each prove req can have at most MaxSwap or MaxPoolNum whichever hits first
func (c *OneChain) ProcessReceipts(receipts []*types.Receipt, sender common.Address) ([]*OneProveReq, error) {
	rows, _ := c.db.Pools(context.Background(), c.ChainID)
	poolidMap := make(map[common.Hash]binding.PoolKey)
	for _, row := range rows {
		poolidMap[Hex2hash(row.Poolid)] = row.Poolkey
	}

	pmAddr := Hex2addr(c.PoolMgr)
	swapEvId := Hex2hash(SwapEvId)
	claimev, err := c.db.ClaimerGet(context.Background(), dal.ClaimerGetParams{
		Chid:   c.ChainID,
		Router: Addr2hex(sender),
	})
	found, _ := dal.ChkQueryRow(err)
	if !found {
		return nil, fmt.Errorf("please contact Brevis team for proper setup of %s", sender)
	}

	// go over tx receipts to filter eligible Swap logs
	var logs []OneLog
	for _, r := range receipts {
		// all logs in one receipt has same logIdxOffset
		logIdxOffset := r.Logs[0].Index
		hasAppend := false // for this receipt, whether we have appended OneLog to logs, if true we need to set last OneLog TxGasCap
		for _, l := range r.Logs {
			if l.Address == pmAddr && l.Topics[0] == swapEvId {
				poolid := l.Topics[1]
				if _, ok := poolidMap[poolid]; !ok {
					// skip ineligible poolids
					continue
				}
				// skip if sender doesn't match
				if sender != common.BytesToAddress(l.Topics[2][12:]) {
					continue
				}
				// append
				hasAppend = true
				logs = append(logs, OneLog{
					Log:          l,
					LogIdxOffset: logIdxOffset,
					// TxGasCap default 0
				})
			}
		}
		if hasAppend { // set last log.TxGasCap of this receipt
			logs[len(logs)-1].TxGasCap = uint32(r.GasUsed * 80 / 100) // actual gas * 0.8
		}
	}

	if len(logs) == 0 {
		return nil, fmt.Errorf("no eligible swaps for sender %s", sender)
	}

	// Circuit supports up to MaxSwaps and MaxPoolNum
	// so we need to split into multiple requests if logs has more. Note we have to keep all swaps exactly ordered as onchain de-dup is by blknum
	// all swaps happen in same blk must be in one batch, so if MaxPool is 32 and within one block there are swaps touching more than 32 pools
	// 33rd pool swaps will have no effect as circuit poolid check will return 0

	// go over logs, and keep track of unique poolids and count, if exceeds MaxSwaps or MaxPoolNum, creates NewOneProveReq
	var proveReqs []*OneProveReq
	curReq := c.NewOneProveReq(&claimev.Raw)
	proveReqs = append(proveReqs, curReq)

	curPoolMap := make(PoolIdMap) // unique poolids in current req
	// blkNums is sorted ascending
	blkNums, blk2swaps := SwapsByBlock(logs)
	for _, blknum := range blkNums {
		one := blk2swaps[blknum]
		// single block is too big, fail for now and guide requester to use another flow
		if len(one.PoolIds) > circuit.MaxPoolNum || len(one.Logs) > circuit.MaxSwapNum {
			return nil, fmt.Errorf("swaps on block %d exceed limit. pools %d, swaps %d", blknum, len(one.PoolIds), len(one.Logs))
		}
		// swaps includes logs and map of poolid, if within limit, add to curReq
		if len(curReq.Logs)+len(one.Logs) <= circuit.MaxSwapNum &&
			curPoolMap.CombineCount(one.PoolIds) <= circuit.MaxPoolNum {
			curReq.Logs = append(curReq.Logs, one.Logs...)
			curPoolMap.Merge(one.PoolIds)
		} else {
			// can't fit into current req, need to create new req, but first populate req.PoolKey
			for k := range curPoolMap {
				curReq.PoolKey = append(curReq.PoolKey, poolidMap[k])
			}

			curReq = c.NewOneProveReq(&claimev.Raw)
			clear(curPoolMap)
			curReq.Logs = append(curReq.Logs, one.Logs...)
			curPoolMap.Merge(one.PoolIds)
		}
	}
	if len(curReq.PoolKey) == 0 {
		for k := range curPoolMap {
			curReq.PoolKey = append(curReq.PoolKey, poolidMap[k])
		}
	}
	return proveReqs, nil
}

// ReqId is set by server.go
func (c *OneChain) NewOneProveReq(claimev *types.Log) *OneProveReq {
	return &OneProveReq{
		ChainId:    c.ChainID,
		GasPerSwap: c.GasPerSwap,
		GasPerTx:   c.GasPerTx,
		PoolMgr:    c.PoolMgr,
		Logs: []OneLog{{
			Log:          claimev,
			LogIdxOffset: claimev.Index,
		}},
	}
}
