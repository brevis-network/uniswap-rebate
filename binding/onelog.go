package binding

import (
	"database/sql/driver"
	"encoding/json"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type OneLog struct {
	*types.Log   // may also be Claimer event
	LogIdxOffset uint
	TxGasCap     uint32 // tx gas * 0.8. if 0, means next swap is from same tx
}

// only valid for swap, return topics[1]
func (o OneLog) PoolId() common.Hash {
	return o.Topics[1]
}

// all info needed to start proving for one user request, may have multiple app proofs
type ProofInfo struct {
	ReqId   int64
	ChainId uint64
	PoolMgr string
	Logs    []OneLog // fist is claimev, rest are swaps
	//
	GasPerSwap, GasPerTx uint32
}

func (r ProofInfo) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *ProofInfo) Scan(value any) error {
	return json.Unmarshal(value.([]byte), r)
}

type PoolIdMap map[common.Hash]bool

// count all unique keys of both maps
func (m PoolIdMap) CombineCount(a PoolIdMap) int {
	extra := 0
	for k := range a {
		if !m[k] { // not in m, +1
			extra += 1
		}
	}
	return len(m) + extra
}

// add a into m
func (m PoolIdMap) Merge(a PoolIdMap) {
	for k := range a {
		m[k] = true
	}
}

// all swaps in same block
type SameBlkSwaps struct {
	Logs    []OneLog
	PoolIds PoolIdMap
}

// group OneLog by swap blocknum including unique poolids for each group
// also return sorted blknum for ordered iter
func SwapsByBlock(in []OneLog) ([]uint64, map[uint64]SameBlkSwaps) {
	m := make(map[uint64]SameBlkSwaps)
	for _, l := range in {
		ssb := m[l.BlockNumber]
		ssb.Logs = append(ssb.Logs, l)
		if ssb.PoolIds == nil {
			ssb.PoolIds = make(PoolIdMap)
		}
		ssb.PoolIds[l.PoolId()] = true
	}
	keys := make([]uint64, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys, m
}
