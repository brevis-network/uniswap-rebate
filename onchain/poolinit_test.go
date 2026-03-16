package onchain

import (
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestFindPoolInitLogSingleCandidate(t *testing.T) {
	poolMgr := common.HexToAddress("0x00000000000000000000000000000000000000aa")
	wantHash := common.HexToHash("0x1234")
	receipt := &types.Receipt{
		Logs: []*types.Log{
			{
				Address: poolMgr,
				Topics:  []common.Hash{Hex2hash(PoolInitEvId), wantHash},
				Index:   7,
			},
		},
	}

	got, err := FindPoolInitLog(receipt, &poolMgr, nil)
	if err != nil {
		t.Fatalf("FindPoolInitLog err: %v", err)
	}
	if got.Index != 7 {
		t.Fatalf("unexpected log index %d", got.Index)
	}
	if got.Topics[1] != wantHash {
		t.Fatalf("unexpected topic %s", got.Topics[1].Hex())
	}
}

func TestFindPoolInitLogRequiresDisambiguation(t *testing.T) {
	poolMgr := common.HexToAddress("0x00000000000000000000000000000000000000aa")
	receipt := &types.Receipt{
		Logs: []*types.Log{
			{
				Address: poolMgr,
				Topics:  []common.Hash{Hex2hash(PoolInitEvId)},
				Index:   3,
			},
			{
				Address: poolMgr,
				Topics:  []common.Hash{Hex2hash(PoolInitEvId)},
				Index:   5,
			},
		},
	}

	_, err := FindPoolInitLog(receipt, &poolMgr, nil)
	if !errors.Is(err, ErrMultiplePoolInitLog) {
		t.Fatalf("expected ErrMultiplePoolInitLog, got %v", err)
	}

	logIndex := uint(5)
	got, err := FindPoolInitLog(receipt, &poolMgr, &logIndex)
	if err != nil {
		t.Fatalf("FindPoolInitLog with log index err: %v", err)
	}
	if got.Index != logIndex {
		t.Fatalf("unexpected log index %d", got.Index)
	}
}

func TestFindPoolInitLogFiltersPoolManager(t *testing.T) {
	poolMgr := common.HexToAddress("0x00000000000000000000000000000000000000aa")
	other := common.HexToAddress("0x00000000000000000000000000000000000000bb")
	receipt := &types.Receipt{
		Logs: []*types.Log{
			{
				Address: other,
				Topics:  []common.Hash{Hex2hash(PoolInitEvId)},
				Index:   1,
			},
		},
	}

	_, err := FindPoolInitLog(receipt, &poolMgr, nil)
	if !errors.Is(err, ErrNoPoolInitLog) {
		t.Fatalf("expected ErrNoPoolInitLog, got %v", err)
	}
}
