package onchain

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/brevis-network/uniswap-rebate/binding"
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
)

var (
	ZeroAddr common.Address
)

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

type OneChain struct {
	*OneChainConfig
	ec  *ethclient.Client
	mon *mon2.Monitor
	db  *dal.DAL
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
