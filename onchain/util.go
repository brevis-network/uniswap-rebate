package onchain

import (
	"sort"

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

// group OneLog by swap blocknum also return unique poolids for each group
func GroupSwapsByBlock(in []OneLog) (ret []SameBlkSwaps) {
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
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, k := range keys {
		ret = append(ret, m[k])
	}
	return ret
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
