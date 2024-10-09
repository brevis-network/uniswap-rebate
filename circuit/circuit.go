package circuit

import (
	"encoding/hex"

	"github.com/brevis-network/brevis-sdk/sdk"
)

const (
	MaxReceipts = 32
)

var (
	EventIdSwap = sdk.ParseEventID(Hex2Bytes("0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f"))
)

type GasCircuit struct {
	PoolMgr sdk.Uint248 // PoolManager addr
	Sender  sdk.Uint248 // msg.sender of swaps
	Oracle  sdk.Uint248 // price oracle addr, ratio at slot0, value is ratio*10^18. note this is uni to eth ratio so we divide it to convert eth to uni
	PoolId  sdk.Bytes32
	// gas amount of one swap event
	GasPerSwap sdk.Uint248
}

// 1 receipt corresponds to 1 slot
func (c *GasCircuit) Allocate() (maxReceipts, maxStorage, maxTransactions int) {
	return MaxReceipts, MaxReceipts, 0
}

// one receipt has 2 fields, which are same swap log different fields (poolid, sender)
// receipts must be ordered by block num and if multiple receipts have same block num,
// only first corresponding slot has actual data, rests are dummy
// receipts: b1r1, b1r2, b1r3, b2r1
// slots: b1s1, 0, 0, b2s1
func (c *GasCircuit) Define(api *sdk.CircuitAPI, in sdk.DataInput) error {
	api.AssertInputsAreUnique()
	receipts := sdk.NewDataStream(api, in.Receipts)
	// for each receipt, make sure it's from expected pool and msg.sender
	sdk.AssertEach(receipts, func(r sdk.Receipt) sdk.Uint248 {
		swapLog := r.Fields[0]
		swapLog2 := r.Fields[1]
		return api.Uint248.And(
			api.Uint248.IsEqual(swapLog.Contract, c.PoolMgr),
			api.Uint248.IsEqual(swapLog2.Contract, c.PoolMgr),
			api.Uint248.IsEqual(swapLog.EventID, EventIdSwap),
			api.Uint248.IsEqual(swapLog2.EventID, EventIdSwap),
			api.Bytes32.IsEqual(swapLog.Value, c.PoolId),
			api.Uint248.IsEqual(api.ToUint248(swapLog2.Value), c.Sender),
		)
	})

	blockNums := sdk.Map(receipts, func(cur sdk.Receipt) sdk.Uint248 { return api.ToUint248(cur.BlockNum) })
	minBlockNum := sdk.Min(blockNums)
	maxBlockNum := sdk.Max(blockNums)

	// for each swap, eth cost is GasPerSwap*BaseFee, then convert to uni
	totalUni := sdk.ConstUint248(0)
	lastRatio := sdk.ConstUint248(0)
	for i := 0; i < len(in.Receipts.Raw); i++ {
		r := in.Receipts.Raw[i]
		eth := api.Uint248.Mul(r.BlockBaseFee, c.GasPerSwap)
		slot := in.StorageSlots.Raw[i]
		// if slot blocknum isn't 0, receipt blocknum shoould equal slot blocknum
		api.Uint32.AssertIsEqual(r.BlockNum, api.Uint32.Select(
			api.Uint32.IsZero(slot.BlockNum),
			r.BlockNum,
			slot.BlockNum,
		))
		// if slot.BlockNum is 0, use last ratio
		lastRatio = api.Uint248.Select(
			api.Uint248.IsZero(sdk.Uint248(slot.BlockNum)),
			lastRatio,
			api.ToUint248(slot.Value),
		)
		// ratio value is actual ratio * 10^18, this is uni to eth eg. 0.003, so eth / ratio get uni
		eth = api.Uint248.Mul(eth, sdk.ConstUint248(1e18))
		uni, _ := api.Uint248.Div(eth, lastRatio)
		totalUni = api.Uint248.Add(totalUni, uni)
	}

	api.OutputAddress(c.Sender)
	api.OutputBytes32(c.PoolId)
	api.OutputUint(64, minBlockNum)
	api.OutputUint(64, maxBlockNum)
	api.OutputUint(128, totalUni)

	return nil
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
