package circuit

import (
	"encoding/hex"
	"math"

	"github.com/brevis-network/brevis-sdk/sdk"
)

const (
	MaxPerPool  = 32
	MaxPoolNum  = 2
	MaxReceipts = MaxPerPool * MaxPoolNum

	// needed to compare min blknum
	// maxU64 uint64 = math.MaxUint64
	maxU32 uint32 = math.MaxUint32
)

var (
	EventIdSwap = sdk.ParseEventID(Hex2Bytes("0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f"))
	zeroB32     = sdk.ConstBytes32([]byte{0})
)

type GasCircuit struct {
	PoolMgr sdk.Uint248 // PoolManager addr
	Sender  sdk.Uint248 // msg.sender of swaps
	Oracle  sdk.Uint248 // price oracle addr, ratio at slot0, value is ratio*10^18. note this is uni to eth ratio so we divide it to convert eth to uni
	PoolId  [MaxPoolNum]sdk.Bytes32
	// gas amount of one swap event
	GasPerSwap sdk.Uint248
}

// duplicated slots will be empty/dummy
func (c *GasCircuit) Allocate() (maxReceipts, maxStorage, maxTransactions int) {
	return MaxReceipts, MaxReceipts, 0
}

// one receipt has 2 fields, which are same swap log different fields (poolid, sender)
// receipts must be ordered by block num and if multiple receipts have same block num,
// only first corresponding slot has actual data, rests are dummy
// receipts: b1r1, b1r2, b1r3, b2r1
// slots: b1s1, 0, 0, b2s1
// first MaxPerPool data is for poolid[0], etc.
func (c *GasCircuit) Define(api *sdk.CircuitAPI, in sdk.DataInput) error {
	api.AssertInputsAreUnique()
	receipts := sdk.NewDataStream(api, in.Receipts)
	// for each receipt, make sure it's from expected msg.sender
	sdk.AssertEach(receipts, func(r sdk.Receipt) sdk.Uint248 {
		swapLog := r.Fields[0]
		swapLog2 := r.Fields[1]
		return api.Uint248.And(
			api.Uint248.IsEqual(swapLog.Contract, c.PoolMgr),
			api.Uint248.IsEqual(swapLog2.Contract, c.PoolMgr),
			api.Uint248.IsEqual(swapLog.EventID, EventIdSwap),
			api.Uint248.IsEqual(swapLog2.EventID, EventIdSwap),
			api.Uint248.IsEqual(api.ToUint248(swapLog2.Value), c.Sender),
		)
	})

	// per pool min/max blknum and totalUni
	minBlk := [MaxPoolNum]sdk.Uint32{}
	maxBlk := [MaxPoolNum]sdk.Uint32{}
	totalUni := [MaxPoolNum]sdk.Uint248{}
	for i := 0; i < MaxPoolNum; i++ {
		minBlk[i] = sdk.ConstUint32(maxU32)
		maxBlk[i] = sdk.ConstUint32(0)
		totalUni[i] = sdk.ConstUint248(0)
	}
	api.OutputAddress(c.Sender)
	lastRatio := sdk.ConstUint248(0) // save last slot for same block
	// note if lastRatio is per poolid segment, empty poolid will cause div by zero
	// split input into segments
	for poolidx := 0; poolidx < MaxPoolNum; poolidx++ {
		// for each swap, eth cost is GasPerSwap*BaseFee, then convert to uni

		baseIdx := poolidx * MaxPerPool
		// receipt and storage index
		for i := baseIdx; i < MaxPerPool+baseIdx; i++ {
			r := in.Receipts.Raw[i]
			// if r.BlockNum is 0, consider data is a dummy so do nothing
			isDummy := api.Uint32.IsZero(r.BlockNum)
			// ensure poolid matches
			api.Bytes32.AssertIsEqual(c.PoolId[poolidx], api.Bytes32.Select(
				api.ToUint248(isDummy),
				c.PoolId[poolidx],
				r.Fields[0].Value,
			))

			slot := in.StorageSlots.Raw[i]
			// if slot blocknum isn't 0, receipt blocknum should equal slot blocknum
			api.Uint32.AssertIsEqual(r.BlockNum, api.Uint32.Select(
				api.Uint32.Or(isDummy, api.Uint32.IsZero(slot.BlockNum)),
				r.BlockNum,
				slot.BlockNum,
			))
			// if slot.BlockNum is 0, use last ratio
			lastRatio = api.Uint248.Select(
				api.Uint248.Or(api.ToUint248(isDummy), api.Uint248.IsZero(api.ToUint248(slot.BlockNum))),
				lastRatio,
				api.ToUint248(slot.Value),
			)
			eth := api.Uint248.Mul(r.BlockBaseFee, c.GasPerSwap)
			// ratio value is actual ratio * 10^18, this is uni to eth eg. 0.003, so eth / ratio get uni
			eth = api.Uint248.Mul(eth, sdk.ConstUint248(1e18))
			uni, _ := api.Uint248.Div(eth, lastRatio)
			totalUni[poolidx] = api.Uint248.Add(totalUni[poolidx], uni)
			// possible: receipt[0] for min, last valid receipt for max?
			minBlk[poolidx] = api.Uint32.Select(
				// not dummy, and receipt has smaller blocknum
				api.Uint32.And(api.Uint32.Not(isDummy), api.Uint32.IsLessThan(r.BlockNum, minBlk[poolidx])),
				r.BlockNum,
				minBlk[poolidx],
			)
			maxBlk[poolidx] = api.Uint32.Select(
				api.Uint32.And(api.Uint32.Not(isDummy), api.Uint32.IsGreaterThan(r.BlockNum, maxBlk[poolidx])),
				r.BlockNum,
				maxBlk[poolidx],
			)
		}
		api.OutputBytes32(c.PoolId[poolidx])
		api.OutputUint32(32, sdk.ConstUint32(0)) // fill 0 as contract expects 8 bytes blknum
		api.OutputUint32(32, minBlk[poolidx])
		api.OutputUint32(32, sdk.ConstUint32(0)) // fill 0 as contract expects 8 bytes blknum
		api.OutputUint32(32, maxBlk[poolidx])
		api.OutputUint(128, totalUni[poolidx])
	}
	return nil
}

func DefaultCircuit() *GasCircuit {
	ret := &GasCircuit{
		PoolMgr:    sdk.ConstUint248(0),
		Sender:     sdk.ConstUint248(0),
		Oracle:     sdk.ConstUint248(0),
		GasPerSwap: sdk.ConstUint248(0),
	}
	for i := 0; i < MaxPoolNum; i++ {
		ret.PoolId[i] = zeroB32
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
