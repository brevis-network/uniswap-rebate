package circuit

import (
	"math/big"
	"testing"

	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/brevis-network/brevis-sdk/test"

	"github.com/ethereum/go-ethereum/common"
)

const (
	PoolMgr  = "0x1111"
	Sender   = "0x2222"
	Oracle   = "0x3333"
	PoolId   = "0x4444"
	SwapEvId = "0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f"
)

func TestCircuit(t *testing.T) {
	app, _ := sdk.NewBrevisApp(1)
	// ========== receipts
	// first 2 receipts from same block
	app.AddReceipt(
		sdk.ReceiptData{
			BlockNum:     big.NewInt(1),
			BlockBaseFee: big.NewInt(1e9),
			Fields: [4]sdk.LogFieldData{
				newLog(1, 1, PoolId),
				newLog(1, 2, Sender),
			},
		},
	)
	app.AddReceipt(
		sdk.ReceiptData{
			BlockNum:     big.NewInt(1),
			BlockBaseFee: big.NewInt(1e9),
			Fields: [4]sdk.LogFieldData{
				newLog(2, 1, PoolId),
				newLog(2, 2, Sender),
			},
		},
	)
	app.AddReceipt(
		sdk.ReceiptData{
			BlockNum:     big.NewInt(2),
			BlockBaseFee: big.NewInt(2e9),
			Fields: [4]sdk.LogFieldData{
				newLog(1, 1, PoolId),
				newLog(1, 2, Sender),
			},
		},
	)
	// ========== slots
	app.AddStorage(
		sdk.StorageData{
			BlockNum:     big.NewInt(1),
			BlockBaseFee: big.NewInt(1e9),
			Address:      common.HexToAddress(Oracle),
			Slot:         common.HexToHash("0x00"),
			Value:        common.BigToHash(big.NewInt(1e16)),
		}, // first slot, no need for index
	)
	// storage 1 is dummy
	app.AddStorage(
		sdk.StorageData{
			BlockNum:     big.NewInt(2),
			BlockBaseFee: big.NewInt(2e9),
			Address:      common.HexToAddress(Oracle),
			Slot:         common.HexToHash("0x00"),
			Value:        common.BigToHash(big.NewInt(2e16)),
		},
		2, // specify index 2 so index 1 is dummy
	)

	defaultC := DefaultCircuit()
	newC := NewCircuit()
	circuitInput, err := app.BuildCircuitInput(newC)
	check(err)
	test.IsSolved(t, defaultC, newC, circuitInput)

	test.ProverSucceeded(t, defaultC, newC, circuitInput)
}

var (
	ContractAddr = common.HexToAddress(PoolMgr)
	EventId      = common.HexToHash(SwapEvId)
)

func newLog(logIdx, fieldIdx uint, value string) sdk.LogFieldData {
	return sdk.LogFieldData{
		Contract:   ContractAddr,
		LogPos:     logIdx,
		EventID:    EventId,
		IsTopic:    true,
		FieldIndex: fieldIdx,
		Value:      common.HexToHash(value),
	}
}

func NewCircuit() *GasCircuit {
	ret := &GasCircuit{
		PoolMgr:    sdk.ConstUint248(Hex2Bytes(PoolMgr)),
		Sender:     sdk.ConstUint248(Hex2Bytes(Sender)),
		Oracle:     sdk.ConstUint248(Hex2Bytes(Oracle)),
		GasPerSwap: sdk.ConstUint248(50000),
	}
	ret.PoolId[0] = sdk.ConstBytes32(Hex2Bytes(PoolId))
	for i := 1; i < MaxPoolNum; i++ {
		ret.PoolId[i] = zeroB32
	}
	return ret
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
