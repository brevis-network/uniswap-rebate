package circuit

import (
	"encoding/hex"
	"math"

	"github.com/brevis-network/brevis-sdk/sdk"
)

const (
	MaxPoolNum  = 4
	MaxSwapNum  = 1023 // need 1 more for claimer
	MaxReceipts = MaxSwapNum + 1

	maxU32 uint32 = math.MaxUint32
)

var (
	// single event that tells us claimer address, event is Claimer(address)
	EventIdClaimer = sdk.ParseEventID(Hex2Bytes("0x8d5763b8f1aa10a2a7039efd8f390755df967d1e68f0a76dd56ceee013227162"))
	// all swaps of same router
	EventIdSwap = sdk.ParseEventID(Hex2Bytes("0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f"))
	// const
	zeroB32 = sdk.ConstBytes32([]byte{0})
)

type GasCircuit struct {
	PoolMgr sdk.Uint248                 // PoolManager addr
	Sender  sdk.Uint248                 // msg.sender of swaps
	PoolKey [MaxPoolNum * 5]sdk.Bytes32 // each poolkey has 5 fields, poolid = keccak(abi.encode(poolkey))
	// gas rebate of one swap event and per tx.
	// rebate gas for one tx is `n * (rebatePerSwap+rebatePerHook) + rebateFixed` where n is number of valid swaps in this tx.
	GasPerSwap, GasPerTx sdk.Uint248
	// compare the result computed by `n * (rebatePerSwap+rebatePerHook) + rebateFixed` with tx's actual gas usage * 0.8  and choose the smaller one.
	// so here the value is tx actual gas * 0.8. Note if a tx has k valid swaps, there are k entries and first k-1 are all 0, only last one is actual gas cap
	TxGasCap [MaxSwapNum]sdk.Uint248
}

func (c *GasCircuit) Allocate() (maxReceipts, maxStorage, maxTransactions int) {
	return MaxSwapNum + 1, 0, 0
}

// receipt[0] is claimer event, [1:] are all swaps
// one swap receipt has 2 fields, poolid and sender from the same swap log
func (c *GasCircuit) Define(api *sdk.CircuitAPI, in sdk.DataInput) error {
	// each receipt must be unique
	api.AssertInputsAreUnique()

	// check pool has non-zero hook and compute poolID, solidity in memory struct doesn't do packing so each field occupies 32bytes
	var poolIDs [MaxPoolNum]sdk.Bytes32
	for i := 0; i < MaxPoolNum; i++ {
		// check hook isn't zero
		idx := i * 5 // idx of first slot (of total 5)
		api.Uint248.AssertIsDifferent(api.ToUint248(c.PoolKey[idx+4]), sdk.ConstUint248(0))
		poolIDs[i] = api.Keccak256(c.PoolKey[idx:idx+5], []int32{256, 256, 256, 256, 256})
	}

	// for each receipt, ensure its poolid at least matches one in poolIDs
	receipts := sdk.NewDataStream(api, in.Receipts)
	sdk.AssertEach(receipts, func(r sdk.Receipt) sdk.Uint248 {
		isSwap := api.Uint248.And(
			api.Uint248.IsEqual(r.Fields[0].Contract, c.PoolMgr),
			api.Uint248.IsEqual(r.Fields[1].Contract, c.PoolMgr),
			api.Uint248.IsEqual(r.Fields[0].EventID, EventIdSwap),
			api.Uint248.IsEqual(r.Fields[1].EventID, EventIdSwap),
			api.Uint248.IsEqual(api.ToUint248(r.Fields[1].Value), c.Sender),
		)
		eligible := sdk.ConstUint32(0) // check event poolid with all poolids
		for j := 0; j < MaxPoolNum; j++ {
			// if event poolid matches poolid, set eligiblePoolId to 1, otherwise keep as is
			eligible = api.Uint32.Select(
				sdk.Uint32{Val: api.Bytes32.IsEqual(poolIDs[j], r.Fields[0].Value).Val},
				sdk.ConstUint32(1),
				eligible,
			)
		}
		// if swap, must be eligible, if not swap, must be claim ev
		return api.Uint248.Select(
			isSwap,
			api.ToUint248(eligible),
			api.Uint248.And(
				api.Uint248.IsEqual(r.Fields[0].Contract, c.Sender),
				api.Uint248.IsEqual(r.Fields[0].EventID, EventIdClaimer),
			),
		)
	})

	curTxGas := sdk.ConstUint248(0)    // sum gas of the same tx
	totalRebate := sdk.ConstUint248(0) // output, sum of rebate gas * gas price

	for i := 0; i < MaxSwapNum; i++ {
		r := in.Receipts.Raw[i+1] // first receipt is claimer so swap starts from 1 but TxGasCap starts from 0
		// add swap to curTxGas
		curTxGas = api.Uint248.Add(curTxGas, c.GasPerSwap)
		// now check TxGasCap, if 0, means more receipts belong to same tx
		// if not 0, means last swap of tx, also add GasPerTx, and compare to TxGasCap, adds smaller one to totalRebate
		lastSwap := api.Uint248.IsGreaterThan(c.TxGasCap[i], sdk.ConstUint248(0))
		curTxGas = api.Uint248.Select(
			lastSwap,
			api.Uint248.Add(curTxGas, c.GasPerTx), // add fixed per tx gas
			curTxGas,
		)
		// no need to check lastSwap because if c.TxGasCap[i] is 0, toAdd is guaranteed to also be 0
		toAdd := api.Uint248.Select(
			api.Uint248.IsLessThan(curTxGas, c.TxGasCap[i]),
			curTxGas,
			c.TxGasCap[i],
		)
		// multiple gas by gas fee and add to total, if not last swap, toAdd is 0 so no change
		// Note for dummy receipts, corresponding TxGasCap should all be 0 so toAdd is also 0
		totalRebate = api.Uint248.Add(totalRebate, api.Uint248.Mul(toAdd, r.BlockBaseFee))
		// if lastSwap, reset curTxGas to 0 for new tx next
		curTxGas = api.Uint248.Select(
			lastSwap,
			sdk.ConstUint248(0),
			curTxGas,
		)
	}
	// if we can make receipts only include swap, it's easy but if it also has claim ev, have to do it manually
	//blkNums := sdk.Map(receipts, func(r sdk.Receipt) sdk.Uint32 {})
	minBlk := sdk.Reduce(receipts, sdk.ConstUint32(maxU32), func(minBlk sdk.Uint32, r sdk.Receipt) sdk.Uint32 {
		isSwap := sdk.Uint32{Val: api.Uint248.IsEqual(r.Fields[0].EventID, EventIdSwap).Val}
		return api.Uint32.Select(
			api.Uint32.And(isSwap, api.Uint32.IsLessThan(r.BlockNum, minBlk)),
			r.BlockNum,
			minBlk,
		)
	})
	maxBlk := sdk.Reduce(receipts, sdk.ConstUint32(0), func(maxBlk sdk.Uint32, r sdk.Receipt) sdk.Uint32 {
		isSwap := sdk.Uint32{Val: api.Uint248.IsEqual(r.Fields[0].EventID, EventIdSwap).Val}
		return api.Uint32.Select(
			api.Uint32.And(isSwap, api.Uint32.IsGreaterThan(r.BlockNum, maxBlk)),
			r.BlockNum,
			minBlk,
		)
	})

	// output router and claimer address
	api.OutputAddress(c.Sender)
	api.OutputAddress(api.ToUint248(in.Receipts.Raw[0].Fields[0].Value))

	api.OutputUint32(32, sdk.ConstUint32(0)) // fill 0 as contract expects 8 bytes blknum
	api.OutputUint32(32, minBlk)
	api.OutputUint32(32, sdk.ConstUint32(0)) // fill 0 as contract expects 8 bytes blknum
	api.OutputUint32(32, maxBlk)
	api.OutputUint(128, totalRebate)
	return nil
}

func DefaultCircuit() *GasCircuit {
	ret := &GasCircuit{
		PoolMgr:    sdk.ConstUint248(0),
		Sender:     sdk.ConstUint248(0),
		GasPerSwap: sdk.ConstUint248(0),
		GasPerTx:   sdk.ConstUint248(0),
	}
	for i := 0; i < MaxPoolNum*5; i++ {
		ret.PoolKey[i] = zeroB32
	}
	for i := 0; i < MaxSwapNum; i++ {
		ret.TxGasCap[i] = sdk.ConstUint248(0)
	}
	return ret
}

// ===== utils =====
func Hex2Bytes(s string) (b []byte) {
	if len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
		s = s[2:]
	}
	// hex.DecodeString expects an even-length string
	if len(s)%2 == 1 {
		s = "0" + s
	}
	b, _ = hex.DecodeString(s)
	return b
}
