package onchain

import (
	"context"

	"github.com/brevis-network/uniswap-rebate/binding"
	"github.com/brevis-network/uniswap-rebate/circuit"
	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// evlog is full event not types.Log because we need to add Value/Scan sql interface
func (c *OneChain) MonClaimer() {
	helper := Hex2addr(circuit.ClaimHelp)
	filter, _ := binding.NewClaimHelpFilterer(helper, c.ec)
	go c.mon.MonAddr(mon2.PerAddrCfg{
		Addr:    helper,
		ChkIntv: GetLogIntv,
		AbiStr:  binding.ClaimHelpMetaData.ABI,
	}, func(s string, l types.Log) {
		if s != "Claimer" {
			log.Error("unexpected ev:", s)
			return
		}
		ev, err := filter.ParseClaimer(l)
		if err != nil {
			log.Error("parse log err: ", err)
			return
		}
		log.Infoln("router:", ev.Router, "claimer:", ev.Claimer)
		// save into db for later use
		c.db.ClaimerAdd(context.Background(), dal.ClaimerAddParams{
			Chid:   c.ChainID,
			Router: Addr2hex(ev.Router),
			Evlog:  *ev,
		})
	})
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
			Chid:    c.ChainID,
			Poolid:  poolid,
			Poolkey: poolK,
		})
		if err != nil {
			log.Errorln("pooladd err:", err)
		}
	})
}
