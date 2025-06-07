package onchain

import (
	"fmt"

	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/brevis-network/uniswap-rebate/binding"
	"github.com/brevis-network/uniswap-rebate/circuit"
	"github.com/celer-network/goutils/log"
)

// block basefee and timestamp
type OneBlock struct {
	BaseFee, Timestamp uint64
}

// all needed for one circuit proof, supports multi pools
type OneProveReq struct {
	ChainId uint64
	PoolMgr string
	// circuit fields
	GasPerSwap, GasPerTx uint32
	// unique poolkeys from logs
	PoolKey []binding.PoolKey

	// circuit input receipts, first is claimer, rest are swaps
	// all Swap logs have same sender, sorted by blknum
	// swaps from same tx must be together (ok to re-order)
	Logs []binding.OneLog

	// set by server
	ReqId int64
}

// create circuit obj from prove req, used by prover
func (r *OneProveReq) NewCircuit() *circuit.GasCircuit {
	ret := circuit.DefaultCircuit()
	ret.PoolMgr = sdk.ConstUint248(Hex2Bytes(r.PoolMgr))
	ret.GasPerSwap = sdk.ConstUint32(r.GasPerSwap)
	ret.GasPerTx = sdk.ConstUint32(r.GasPerTx)

	if len(r.PoolKey) > circuit.MaxPoolNum {
		// more pools than circuit can hadle
		log.Warn(fmt.Sprintf("req %d has %d pools, more than MaxPoolNum %d", r.ReqId, len(r.PoolKey), circuit.MaxPoolNum))
	}
	for i, poolkey := range r.PoolKey {
		idx := i * 5
		// todo: break if idx > len(ret.PoolKey)
		ret.PoolKey[idx] = sdk.ConstFromBigEndianBytes(poolkey.Currency0[:])
		ret.PoolKey[idx+1] = sdk.ConstFromBigEndianBytes(poolkey.Currency1[:])
		ret.PoolKey[idx+2] = sdk.ConstFromBigEndianBytes(poolkey.Fee.Bytes())
		ret.PoolKey[idx+3] = sdk.ConstFromBigEndianBytes(poolkey.TickSpacing.Bytes())
		ret.PoolKey[idx+4] = sdk.ConstFromBigEndianBytes(poolkey.Hooks[:])
	}
	// skip first Claim ev
	for i, swap := range r.Logs[1:] {
		ret.TxGasCap[i] = sdk.ConstUint32(swap.TxGasCap)
	}
	return ret
}
