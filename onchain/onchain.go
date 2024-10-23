package onchain

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"time"

	"github.com/brevis-network/uniswap-rebate/binding"
	"github.com/brevis-network/uniswap-rebate/circuit"
	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
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
	Swap         *types.Log
	LogIdxOffset uint
}

type OneBlock struct {
	BaseFee   uint64
	SlotValue [32]byte
}

// all needed for one circuit proof, supports multi pools
type OneProveReq struct {
	ChainId, GasPerSwap     uint64
	PoolMgr, Sender, Oracle string
	// unique poolids from logs
	PoolIds []string
	// all logs have same sender, sorted by blknum
	Logs []OneLog
	// blknum -> info about at this block, include all blocknums from Logs
	Blks map[uint64]OneBlock
}

// sort logs and populate Blks from m
func (r *OneProveReq) Fix(m map[uint64]OneBlock) {
	sort.Slice(r.Logs, func(i, j int) bool {
		return r.Logs[i].Swap.BlockNumber < r.Logs[j].Swap.BlockNumber
	})
	for _, l := range r.Logs {
		// ok to set again as it's same anyway
		r.Blks[l.Swap.BlockNumber] = m[l.Swap.BlockNumber]
	}
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

func (c *OneChain) MonPoolInit() {
	pmAddr := Hex2addr(c.PoolMgr)
	filter, _ := binding.NewPoolMgrFilterer(pmAddr, c.ec)
	go c.mon.MonAddr(mon2.PerAddrCfg{
		Addr:    pmAddr,
		ChkIntv: GetLogIntv,
		AbiStr:  binding.PoolMgrMetaData.ABI,
		Topics:  [][]common.Hash{{Hex2hash(PoolInitEvId)}}, // binpool Initialize event id to reduce log data
	}, func(s string, l types.Log) {
		if s != "Initialize" {
			log.Error("wrong event: ", s)
			return
		}
		initEv, err := filter.ParseInitialize(l)
		if err != nil {
			log.Error("parse log err: ", err)
			return
		}
		// skip zero hook pools
		if initEv.Hooks == ZeroAddr {
			return
		}
		poolid := Hash2Hex(initEv.Id)
		poolK := binding.PoolKey{
			Currency0:   initEv.Currency0,
			Currency1:   initEv.Currency1,
			Fee:         initEv.Fee,
			TickSpacing: initEv.TickSpacing,
			Hooks:       initEv.Hooks,
		}
		log.Infoln("add pool", poolid)
		err = c.db.PoolAdd(context.Background(), dal.PoolAddParams{
			Poolid:  poolid,
			Poolkey: poolK,
		})
		if err != nil {
			log.Errorln("pooladd err:", err)
		}
	})
}

// txlist is hex string of tx hash, return non-nil err if any TransactionReceipt has err
func (c *OneChain) FetchTxReceipts(txlist []string) ([]*types.Receipt, error) {
	var ret []*types.Receipt
	for _, tx := range txlist {
		r, err := c.ec.TransactionReceipt(context.Background(), Hex2hash(tx))
		if err != nil {
			return ret, err
		}
		ret = append(ret, r)
	}
	return ret, nil
}

// go through receipt.Logs, check poolid is in db (eligible w/ non-zero hook addr)
// fetch block basefee and storage, return ready to use start prove reqs
func (c *OneChain) ProcessReceipts(receipts []*types.Receipt) ([]*OneProveReq, error) {
	poolids, _ := c.db.PoolIds(context.Background())
	poolidMap := make(map[[32]byte]bool)
	for _, pid := range poolids {
		poolidMap[Hex2hash(pid)] = true
	}
	// poolid to list of entries
	logByPool := make(map[[32]byte][]OneLog)
	pmAddr := Hex2addr(c.PoolMgr)
	oracle := Hex2addr(c.Oracle)
	swapEvId := Hex2hash(SwapEvId)
	var sender common.Address // default zero addr, will be set from first eligible log

	blkMap := make(map[uint64]OneBlock)
	for _, r := range receipts {
		// all logs in one receipt has same logIdxOffset
		logIdxOffset := r.Logs[0].Index
		for _, l := range r.Logs {
			if l.Address == pmAddr && l.Topics[0] == swapEvId {
				poolid := l.Topics[1]
				if !poolidMap[poolid] {
					// skip ineligible poolids
					continue
				}
				// first eligible log, set sender
				logSender := common.BytesToAddress(l.Topics[2][12:])
				if sender == ZeroAddr {
					sender = logSender
				} else if sender != logSender {
					// skip if sender already set but logSender is different
					continue
				}
				// now append log to correct list
				logByPool[poolid] = append(logByPool[poolid], OneLog{
					Swap:         l,
					LogIdxOffset: logIdxOffset,
				})
				// new blocknum, we could make the map at OneChain level or save to db as info is immutable
				if _, ok := blkMap[l.BlockNumber]; !ok {
					// get basefee and slot
					blkNum := new(big.Int).SetUint64(l.BlockNumber)
					b, err := c.ec.BlockByNumber(context.Background(), blkNum)
					if err != nil {
						log.Errorln("get block", l.BlockNumber, "err:", err)
						continue
					}
					value, err := c.ec.StorageAt(context.Background(), oracle, ZeroHash, blkNum)
					if err != nil {
						log.Errorln("StorageAt", l.BlockNumber, oracle, "err:", err)
						continue
					}
					var slotV common.Hash
					copy(slotV[32-len(value):], value)
					blkMap[l.BlockNumber] = OneBlock{
						BaseFee:   b.BaseFee().Uint64(),
						SlotValue: slotV,
					}
				}
			}
		}
	}

	// OneProveReq supports up to MaxReceipts and MaxPoolNum
	// so we need to split into multiple requests if logByPool has more
	// use simple algo for now
	var proveReqs []*OneProveReq
	for _, b := range SplitMapIntoBatches(logByPool, circuit.MaxPoolNum, circuit.MaxReceipts) {
		// one b has up to limit pools and logs
		req := c.NewOneProveReq(sender)
		for k, v := range b {
			req.PoolIds = append(req.PoolIds, Hash2Hex(k))
			req.Logs = append(req.Logs, v...)
		}
		req.Fix(blkMap)
		proveReqs = append(proveReqs, req)
	}
	return proveReqs, nil
}

func (c *OneChain) NewOneProveReq(sender common.Address, pools ...string) *OneProveReq {
	return &OneProveReq{
		ChainId:    c.ChainID,
		GasPerSwap: c.GasPerSwap,
		PoolMgr:    c.PoolMgr,
		Sender:     Addr2hex(sender),
		Oracle:     c.Oracle,
		PoolIds:    pools,
		Blks:       make(map[uint64]OneBlock),
	}
}

func Hex2addr(addr string) common.Address {
	return common.HexToAddress(addr)
}

// 0x prefix, only hex, all lower case
func Addr2hex(addr common.Address) string {
	return "0x" + hex.EncodeToString(addr[:])
}

func Hex2hash(hexstr string) common.Hash {
	return common.HexToHash(hexstr)
}

func Hash2Hex(h [32]byte) string {
	return "0x" + hex.EncodeToString(h[:])
}
