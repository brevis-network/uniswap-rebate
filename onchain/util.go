package onchain

import (
	"encoding/hex"
	"slices"

	"github.com/ethereum/go-ethereum/common"
)

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

func Hex2addr(addr string) common.Address {
	return common.HexToAddress(addr)
}

// 0x prefix, only hex, all lower case
func Addr2hex(addr common.Address) string {
	return "0x" + hex.EncodeToString(addr[:])
}

func Hex2hash(hexstr string) common.Hash {
	return common.HexToHash(hexstr)
}

func Hash2Hex(h common.Hash) string {
	return "0x" + hex.EncodeToString(h[:])
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

// SplitMapIntoBatches splits a map[K]V into batches based on maxKeys and maxItems constraints
// where V must be a slice
func SplitMapIntoBatches[K comparable, V ~[]E, E any](input map[K]V, maxKeys, maxItems int) []map[K]V {
	if len(input) == 0 {
		return []map[K]V{}
	}

	result := make([]map[K]V, 0)
	currentBatch := make(map[K]V)
	currentItems := 0

	for k, v := range input {
		// If adding this entry would exceed either constraint,
		// add current(if non empty) to result and start a new batch
		if len(currentBatch) >= maxKeys || currentItems+len(v) > maxItems {
			if len(currentBatch) > 0 {
				result = append(result, currentBatch)
				currentBatch = make(map[K]V)
				currentItems = 0
			}
		}

		// Handle case where a single value exceeds maxItems
		if len(v) > maxItems {
			// Split the slice into smaller chunks
			for i := 0; i < len(v); i += maxItems {
				end := i + maxItems
				if end > len(v) {
					end = len(v)
				}
				// last batch may not be full so next loop iter will add next pool if possible
				singleBatch := map[K]V{k: v[i:end]}
				result = append(result, singleBatch)
			}
			continue
		}

		currentBatch[k] = v
		currentItems += len(v)
	}

	// Add the last batch if it's not empty
	if len(currentBatch) > 0 {
		result = append(result, currentBatch)
	}

	return result
}
