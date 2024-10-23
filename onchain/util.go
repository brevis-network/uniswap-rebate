package onchain

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
