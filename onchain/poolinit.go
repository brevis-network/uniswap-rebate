package onchain

import (
	"context"
	"errors"
	"fmt"

	"github.com/brevis-network/uniswap-rebate/binding"
	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	ErrNoPoolInitLog       = errors.New("no Pool Initialize log found in receipt")
	ErrMultiplePoolInitLog = errors.New("multiple Pool Initialize logs found in receipt")
	ErrZeroHookPool        = errors.New("pool Initialize event has zero hooks")
)

type PoolInitRecord struct {
	Event   *binding.PoolMgrInitialize
	PoolID  string
	PoolKey binding.PoolKey
}

func FindPoolInitLog(receipt *types.Receipt, poolMgr *common.Address, logIndex *uint) (*types.Log, error) {
	if receipt == nil {
		return nil, fmt.Errorf("%w: nil receipt", ErrNoPoolInitLog)
	}

	var candidates []*types.Log
	var candidateIdx []uint
	initEvID := Hex2hash(PoolInitEvId)

	for _, lg := range receipt.Logs {
		if len(lg.Topics) == 0 || lg.Topics[0] != initEvID {
			continue
		}
		if poolMgr != nil && lg.Address != *poolMgr {
			continue
		}
		if logIndex != nil && lg.Index != *logIndex {
			continue
		}
		candidates = append(candidates, lg)
		candidateIdx = append(candidateIdx, lg.Index)
	}

	switch len(candidates) {
	case 0:
		msg := "no matching Initialize log"
		if poolMgr != nil {
			msg += " for pool manager " + poolMgr.Hex()
		}
		if logIndex != nil {
			msg += fmt.Sprintf(" at log index %d", *logIndex)
		}
		return nil, fmt.Errorf("%w: %s", ErrNoPoolInitLog, msg)
	case 1:
		return candidates[0], nil
	default:
		return nil, fmt.Errorf("%w: candidate log indexes %v", ErrMultiplePoolInitLog, candidateIdx)
	}
}

func AddPoolInitFromLog(ctx context.Context, db *dal.DAL, chainID uint64, filter *binding.PoolMgrFilterer, lg types.Log) (*PoolInitRecord, error) {
	if len(lg.Topics) == 0 || lg.Topics[0] != Hex2hash(PoolInitEvId) {
		return nil, fmt.Errorf("log %d is not Pool Initialize", lg.Index)
	}

	initEv, err := filter.ParseInitialize(lg)
	if err != nil {
		return nil, fmt.Errorf("parse Initialize log %d: %w", lg.Index, err)
	}
	if initEv.Hooks == ZeroAddr {
		return nil, ErrZeroHookPool
	}

	poolKey := binding.PoolKey{
		Currency0:   initEv.Currency0,
		Currency1:   initEv.Currency1,
		Fee:         initEv.Fee,
		TickSpacing: initEv.TickSpacing,
		Hooks:       initEv.Hooks,
	}
	poolID := Hash2Hex(initEv.Id)

	err = db.PoolAdd(ctx, dal.PoolAddParams{
		Chid:    chainID,
		Poolid:  poolID,
		Poolkey: poolKey,
	})
	if err != nil {
		return nil, fmt.Errorf("PoolAdd %s: %w", poolID, err)
	}

	return &PoolInitRecord{
		Event:   initEv,
		PoolID:  poolID,
		PoolKey: poolKey,
	}, nil
}
