package binding

import (
	"database/sql/driver"
	"encoding/json"
	"slices"

	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/brevis-network/uniswap-rebate/circuit"
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

// new app circuit from info and swaps
func NewCircuit(info *ProofInfo, swaps []OneLog, poolkeys []PoolKey) *circuit.GasCircuit {
	ret := &circuit.GasCircuit{
		PoolMgr:    sdk.ConstUint248(common.Hex2Bytes(info.PoolMgr)),
		GasPerSwap: sdk.ConstUint32(info.GasPerSwap),
		GasPerTx:   sdk.ConstUint32(info.GasPerTx),
	}
	for i, swap := range swaps {
		ret.TxGasCap[i] = sdk.ConstUint32(swap.TxGasCap)
	}
	for i, poolkey := range poolkeys {
		idx := i * 5
		ret.PoolKey[idx] = sdk.ConstFromBigEndianBytes(poolkey.Currency0[:])
		ret.PoolKey[idx+1] = sdk.ConstFromBigEndianBytes(poolkey.Currency1[:])
		ret.PoolKey[idx+2] = sdk.ConstFromBigEndianBytes(poolkey.Fee.Bytes())
		ret.PoolKey[idx+3] = sdk.ConstFromBigEndianBytes(poolkey.TickSpacing.Bytes())
		ret.PoolKey[idx+4] = sdk.ConstFromBigEndianBytes(poolkey.Hooks[:])
	}
	return ret
}

// all info needed to start proving for one user request, may have multiple app proofs
type ProofInfo struct {
	ReqId   int64
	ChainId uint64 // src chain id
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

// a group of swaps, eg. swaps in same block
type SwapsGroup struct {
	Logs    []OneLog
	PoolIds PoolIdMap
}

// group OneLog by swap blocknum including unique poolids for each group
// also return sorted blknum for ordered iter
func SwapsByBlock(in []OneLog) ([]uint64, map[uint64]*SwapsGroup) {
	m := make(map[uint64]*SwapsGroup)
	for _, l := range in {
		ssb := m[l.BlockNumber]
		if ssb == nil {
			ssb = NewSwapsGroup()
			m[l.BlockNumber] = ssb
		}
		ssb.Logs = append(ssb.Logs, l)
		ssb.PoolIds[l.PoolId()] = true
	}
	keys := make([]uint64, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys, m
}

// empty SwapsGroup
func NewSwapsGroup() *SwapsGroup {
	return &SwapsGroup{
		PoolIds: make(PoolIdMap),
	}
}

// if no exceed max, append a.Logs to s and merge their PoolIdMap, return s, false
// if exceed, return a, true.
func (s *SwapsGroup) Merge(a *SwapsGroup, maxLog, maxPool int) (*SwapsGroup, bool) {
	if len(s.Logs)+len(a.Logs) > maxLog || s.PoolIds.CombineCount(a.PoolIds) > maxPool {
		// over max, create new and return
		return a, true
	}
	s.Logs = append(s.Logs, a.Logs...)
	s.PoolIds.Merge(a.PoolIds)
	return s, false
}

// first group by blockNum, iter blknum ascending, pack into group without exceed limit.
func SplitIntoGroups(in []OneLog, maxLog, maxPool int) (ret []*SwapsGroup) {
	if len(in) == 0 {
		return nil
	}
	curGroup := NewSwapsGroup()
	ret = append(ret, curGroup)
	shouldAppend := false // whether to append curGroup to ret
	blkNums, blk2swaps := SwapsByBlock(in)
	for _, blknum := range blkNums {
		thisBlkSwaps := blk2swaps[blknum]
		curGroup, shouldAppend = curGroup.Merge(thisBlkSwaps, maxLog, maxPool)
		if shouldAppend {
			ret = append(ret, curGroup)
		}
	}
	return
}
