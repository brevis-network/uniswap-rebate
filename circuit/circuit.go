package circuit

import (
	"encoding/hex"

	"github.com/brevis-network/brevis-sdk/sdk"
)

const (
	MaxReceipts = 1000
)

var (
	EventIdSwap = sdk.ParseEventID(Hex2Bytes(""))
)

type GasCircuit struct {
	PoolMgr sdk.Uint248 // PoolManager addr
	Sender  sdk.Uint248 // msg.sender of swaps
	PoolId  sdk.Bytes32 // could change to array for multiple pools in one proof
}

func (c *GasCircuit) Allocate() (maxReceipts, maxStorage, maxTransactions int) {
	return MaxReceipts, MaxReceipts, 0
}

// one receipt has 2 fields, which are same swap log different field (poolid, sender)
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

	blockNums := sdk.Map(receipts, func(cur sdk.Receipt) sdk.Uint248 { return cur.BlockNum })
	minBlockNum := sdk.Min(blockNums)
	maxBlockNum := sdk.Max(blockNums)

	api.OutputAddress(c.Sender)
	api.OutputBytes32(c.PoolId)
	api.OutputUint(64, minBlockNum)
	api.OutputUint(64, maxBlockNum)

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
