package onchain

import (
	"context"
	"errors"

	"github.com/brevis-network/uniswap-rebate/binding"
	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const routerBeneficiaryRegistryABI = `[{"type":"event","name":"BeneficiarySet","inputs":[{"name":"chainId","type":"uint64","indexed":true},{"name":"router","type":"address","indexed":true},{"name":"beneficiary","type":"address","indexed":true}],"anonymous":false}]`

// MonBeneficiarySet monitors RouterBeneficiaryRegistry on Unichain and saves source chain router config.
func (c *OneChain) MonBeneficiarySet(registry common.Address) {
	go c.mon.MonAddr(mon2.PerAddrCfg{
		Addr:    registry,
		ChkIntv: GetLogIntv,
		AbiStr:  routerBeneficiaryRegistryABI,
	}, func(s string, l types.Log) {
		if s != "BeneficiarySet" {
			log.Error("unexpected ev:", s)
			return
		}
		if len(l.Topics) != 4 {
			log.Error("unexpected BeneficiarySet topics len:", len(l.Topics))
			return
		}
		srcChid := l.Topics[1].Big().Uint64()
		router := common.BytesToAddress(l.Topics[2].Bytes()[12:])
		beneficiary := common.BytesToAddress(l.Topics[3].Bytes()[12:])

		log.Infoln("beneficiary set, src chain:", srcChid, "router:", router, "beneficiary:", beneficiary)
		err := c.db.ClaimerAdd(context.Background(), dal.ClaimerAddParams{
			Chid:        srcChid,
			Router:      Addr2hex(router),
			Beneficiary: Addr2hex(beneficiary),
		})
		if err != nil {
			log.Error("ClaimerAdd err:", err)
		}
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
		rec, err := AddPoolInitFromLog(context.Background(), c.db, c.ChainID, filter, l)
		if err != nil {
			if errors.Is(err, ErrZeroHookPool) {
				return
			}
			log.Errorln("pool init add err:", err)
			return
		}
		log.Infoln("add pool", rec.PoolID)
	})
}
