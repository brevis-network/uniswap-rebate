package circuit

import (
	"encoding/hex"
	"math"

	"github.com/brevis-network/brevis-sdk/sdk"
)

const (
	MaxPerPool  = 256
	MaxPoolNum  = 16
	MaxReceipts = MaxPerPool * MaxPoolNum

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
	PoolMgr sdk.Uint248                // PoolManager addr
	Sender  sdk.Uint248                // msg.sender of swaps
	PoolKey [MaxPoolNum][5]sdk.Bytes32 // each poolkey has 5 fields, poolid = keccak(abi.encode(poolkey))
	// gas amount of one swap event
	GasPerSwap sdk.Uint248
}

// duplicated slots will be empty/dummy
func (c *GasCircuit) Allocate() (maxReceipts, maxStorage, maxTransactions int) {
	return MaxReceipts, 0, 0
}

// receipt[0] is claimer event, [1:] are all swaps, [MaxPerPool][MaxPerPool]..  MaxPoolNum segments
// one swap receipt has 2 fields, which are same swap log different fields (poolid, sender)
func (c *GasCircuit) Define(api *sdk.CircuitAPI, in sdk.DataInput) error {
	// each receipt must be unique
	api.AssertInputsAreUnique()
	// check first receipt and output router and claimer address
	claimev := in.Receipts.Raw[0].Fields[0]
	api.Uint248.AssertIsEqual(claimev.Contract, c.Sender)
	api.Uint248.AssertIsEqual(claimev.EventID, EventIdClaimer)
	api.OutputAddress(c.Sender)
	api.OutputAddress(api.ToUint248(claimev.Value))

	// check pool has non-zero hook and compute poolID
	var poolIDs [MaxPoolNum]sdk.Bytes32
	for i := 0; i < MaxPoolNum; i++ {
		// check hook isn't zero
		api.Uint248.AssertIsDifferent(api.ToUint248(c.PoolKey[i][4]), sdk.ConstUint248(0))
		poolIDs[i] = api.Keccak256(c.PoolKey[i][:], []int32{256, 256, 256, 256, 256})
	}

	minBlk := sdk.ConstUint32(maxU32)
	maxBlk := sdk.ConstUint32(0)
	swapGas := sdk.ConstUint248(0) // total number of swaps
	// receipt idx start from 1 as 0 is for claimer
	for i := 1; i < MaxReceipts; i++ {
		r := in.Receipts.Raw[i]
		valid := sdk.Uint248{Val: in.Receipts.Toggles[i]}
		// if valid, check receipt fields are expected
		poolidx := (i - 1) / MaxPerPool
		api.Bytes32.AssertIsEqual(poolIDs[poolidx], api.Bytes32.Select(
			valid,
			r.Fields[0].Value,
			poolIDs[poolidx]))
		api.Uint248.AssertIsEqual(valid, api.Uint248.And(
			api.Uint248.IsEqual(r.Fields[0].Contract, c.PoolMgr),
			api.Uint248.IsEqual(r.Fields[1].Contract, c.PoolMgr),
			api.Uint248.IsEqual(r.Fields[0].EventID, EventIdSwap),
			api.Uint248.IsEqual(r.Fields[1].EventID, EventIdSwap),
			api.Uint248.IsEqual(api.ToUint248(r.Fields[1].Value), c.Sender),
		))
		// if valid,  sum gas cost, update min/max blk
		swapGas = api.Uint248.Select(
			valid,
			api.Uint248.Add(swapGas, api.Uint248.Mul(r.BlockBaseFee, c.GasPerSwap)),
			swapGas,
		)
		validU32 := sdk.Uint32{Val: in.Receipts.Toggles[i]}
		minBlk = api.Uint32.Select(
			api.Uint32.And(validU32, api.Uint32.IsLessThan(r.BlockNum, minBlk)),
			r.BlockNum,
			minBlk,
		)
		maxBlk = api.Uint32.Select(
			api.Uint32.And(validU32, api.Uint32.IsGreaterThan(r.BlockNum, maxBlk)),
			r.BlockNum,
			maxBlk,
		)
	}
	return nil
}

func DefaultCircuit() *GasCircuit {
	ret := &GasCircuit{
		PoolMgr:    sdk.ConstUint248(0),
		Sender:     sdk.ConstUint248(0),
		GasPerSwap: sdk.ConstUint248(0),
	}
	for i := 0; i < MaxPoolNum; i++ {
		ret.PoolKey[i] = [5]sdk.Bytes32{zeroB32, zeroB32, zeroB32, zeroB32, zeroB32}
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
