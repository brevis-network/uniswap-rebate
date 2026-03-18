package circuit

import (
	"encoding/hex"
	"log"
	"math"

	"github.com/brevis-network/brevis-sdk/sdk"
)

const (
	MaxPoolNum  = 1
	MaxReceipts = 32
	MaxSwapNum  = MaxReceipts

	maxU32 uint32 = math.MaxUint32
)

var (
	// all swaps of same router
	EventIdSwap = sdk.ParseEventID(Hex2Bytes("0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f"))
	// const
	zeroB32 = sdk.ConstFromBigEndianBytes([]byte{0})
	// gas price cap to 50G wei
	GasPriceCap = sdk.ConstUint248(50_000_000_000)
)

type GasCircuit struct {
	ChainId sdk.Uint64
	PoolMgr sdk.Uint248                 // PoolManager addr
	PoolKey [MaxPoolNum * 5]sdk.Bytes32 // each poolkey has 5 fields, poolid = keccak(abi.encode(poolkey))
	// gas rebate of one swap event and per tx.
	// rebate gas for one tx is `n * (rebatePerSwap+rebatePerHook) + rebateFixed` where n is number of valid swaps in this tx.
	GasPerSwap, GasPerTx sdk.Uint32
	// compare the result computed by `n * (rebatePerSwap+rebatePerHook) + rebateFixed` with tx's actual gas usage * 0.8  and choose the smaller one.
	// so here the value is tx actual gas * 0.8. Note if a tx has k valid swaps, there are k entries and first k-1 are all 0, only last one is actual gas cap
	TxGasCap [MaxSwapNum]sdk.Uint32
}

func (c *GasCircuit) Allocate() (maxReceipts, maxStorage, maxTransactions int) {
	return MaxReceipts, 0, 0
}

// one swap receipt has 2 fields, poolid and sender from the same swap log
func (c *GasCircuit) Define(api *sdk.CircuitAPI, in sdk.DataInput) error {
	// each receipt must be unique
	api.AssertInputsAreUnique()

	// check pool has non-zero hook and compute poolID, solidity in memory struct doesn't do packing so each field occupies 32bytes
	var poolIDs [MaxPoolNum]sdk.Bytes32
	for i := range MaxPoolNum {
		// check hook isn't zero
		idx := i * 5 // idx of first slot (of total 5)
		// if hook is 0, set poolid to 0, otherwise, hash 5 fields
		poolIDs[i] = api.Bytes32.Select(
			api.Bytes32.IsZero(c.PoolKey[idx+4]),
			zeroB32,
			api.Keccak256(c.PoolKey[idx:idx+5], []int32{256, 256, 256, 256, 256}),
		)
	}

	// sender topic of first swap log
	router := api.ToUint248(in.Receipts.Raw[0].Fields[1].Value)
	log.Println(c.PoolMgr, poolIDs, router)

	// build datastream for all swaps
	swaps := sdk.NewDataStream(api, in.Receipts)
	// swaps := sdk.RangeUnderlying(receipts, 0, MaxReceipts)
	// for each swap, ensure it's expected and eligible
	sdk.AssertEach(swaps, func(r sdk.Receipt) sdk.Uint248 {
		isSwap := api.Uint248.And(
			api.Uint248.IsEqual(r.Fields[0].Contract, c.PoolMgr),
			api.Uint248.IsEqual(r.Fields[1].Contract, c.PoolMgr),
			api.Uint248.IsEqual(r.Fields[0].EventID, EventIdSwap),
			api.Uint248.IsEqual(r.Fields[1].EventID, EventIdSwap),
			api.Uint248.IsEqual(api.ToUint248(r.Fields[1].Value), router),
		)
		eligible := sdk.ConstUint32(0) // check event poolid with all poolids
		for j := range MaxPoolNum {
			// if event poolid matches poolid, set eligible to 1, otherwise keep eligible as is
			eligible = api.Uint32.Or(
				eligible,
				sdk.Uint32{Val: api.Bytes32.IsEqual(poolIDs[j], r.Fields[0].Value).Val},
			)
		}
		log.Println(isSwap, eligible)

		return api.Uint248.And(isSwap, api.ToUint248(eligible))
	})

	// if TxGasCap[i] is 0, check receipt i and i+1 are from same tx. no need to check last receipt.
	for i := 0; i < MaxSwapNum-1; i++ {
		cur := in.Receipts.Raw[i]
		next := in.Receipts.Raw[i+1]
		api.Uint32.AssertIsEqual(api.Uint32.Select(
			api.Uint32.IsZero(c.TxGasCap[i]),
			api.Uint32.And(
				api.Uint32.IsEqual(cur.BlockNum, next.BlockNum),
				api.Uint32.IsEqual(cur.MptKeyPath, next.MptKeyPath)),
			sdk.ConstUint32(1),
		), sdk.ConstUint32(1))
	}

	curTxGas := sdk.ConstUint32(0)     // sum gas of the same tx
	totalRebate := sdk.ConstUint248(0) // output, sum of rebate gas * gas price

	for i := range MaxSwapNum {
		// add swap to curTxGas
		curTxGas = api.Uint32.Add(curTxGas, c.GasPerSwap)
		// now check TxGasCap, if 0, means more receipts belong to same tx
		// if not 0, means last swap of tx, also add GasPerTx, and compare to TxGasCap, adds smaller one to totalRebate
		lastSwap := api.Uint32.IsGreaterThan(c.TxGasCap[i], sdk.ConstUint32(0))
		curTxGas = api.Uint32.Select(
			lastSwap,
			api.Uint32.Add(curTxGas, c.GasPerTx), // add fixed per tx gas
			curTxGas,
		)
		// min(curTxGas, TxGasCap)
		// no need to check lastSwap because if c.TxGasCap[i] is 0, toAdd is guaranteed to also be 0
		toAdd := api.Uint32.Select(
			api.Uint32.IsLessThan(curTxGas, c.TxGasCap[i]),
			curTxGas,
			c.TxGasCap[i],
		)

		r := in.Receipts.Raw[i]
		// multiply gas by block base fee and add to total, if not last swap, toAdd is 0 so no change
		// Note for dummy receipts, corresponding TxGasCap should all be 0 so toAdd is also 0
		// gasPrice is min(actual, gasPriceCap)
		gasPrice := api.Uint248.Select(
			api.Uint248.IsLessThan(r.BlockBaseFee, GasPriceCap),
			r.BlockBaseFee,
			GasPriceCap,
		)
		totalRebate = api.Uint248.Add(totalRebate, api.Uint248.Mul(api.ToUint248(toAdd), gasPrice))
		// if lastSwap, reset curTxGas to 0 for new tx next
		curTxGas = api.Uint32.Select(
			lastSwap,
			sdk.ConstUint32(0),
			curTxGas,
		)
	}
	// we could usd blkNums := sdk.Map then Min/Max but need to convert Uint32 to Uint248
	minBlk := sdk.Reduce(swaps, sdk.ConstUint32(maxU32), func(minBlk sdk.Uint32, r sdk.Receipt) sdk.Uint32 {
		return api.Uint32.Select(
			api.Uint32.IsLessThan(r.BlockNum, minBlk),
			r.BlockNum,
			minBlk,
		)
	})
	maxBlk := sdk.Reduce(swaps, sdk.ConstUint32(0), func(maxBlk sdk.Uint32, r sdk.Receipt) sdk.Uint32 {
		return api.Uint32.Select(
			api.Uint32.IsGreaterThan(r.BlockNum, maxBlk),
			r.BlockNum,
			maxBlk,
		)
	})

	// output router and source chain id
	api.OutputUint64(64, c.ChainId)
	api.OutputAddress(router)

	api.OutputUint32(32, sdk.ConstUint32(0)) // fill 0 as contract expects 8 bytes blknum
	api.OutputUint32(32, minBlk)
	api.OutputUint32(32, sdk.ConstUint32(0)) // fill 0 as contract expects 8 bytes blknum
	api.OutputUint32(32, maxBlk)
	api.OutputUint(128, totalRebate)
	return nil
}

func DefaultCircuit() *GasCircuit {
	ret := &GasCircuit{
		ChainId:    sdk.ConstUint64(0),
		PoolMgr:    sdk.ConstUint248(0),
		GasPerSwap: sdk.ConstUint32(0),
		GasPerTx:   sdk.ConstUint32(0),
	}
	for i := range MaxPoolNum * 5 {
		ret.PoolKey[i] = zeroB32
	}
	for i := range MaxSwapNum {
		ret.TxGasCap[i] = sdk.ConstUint32(0)
	}
	return ret
}

// ===== utils =====
func Hex2Bytes(s string) []byte {
	if len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
		s = s[2:]
	}
	// hex.DecodeString expects an even-length string
	if len(s)%2 == 1 {
		s = "0" + s
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}
