package findScore

import (
	"sort"
)

func findScore(nums []int) int64 {
	// Create a slice to keep track of marked indices.
	marked := make([]bool, len(nums))

	// Pair each number with its index and sort by value, then by index.
	type Pair struct {
		Value, Index int
	}

	pairs := make([]Pair, len(nums))
	for i, val := range nums {
		pairs[i] = Pair{Value: val, Index: i}
	}

	// Sort pairs by Value first, then by Index.
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Value == pairs[j].Value {
			return pairs[i].Index < pairs[j].Index
		}
		return pairs[i].Value < pairs[j].Value
	})

	var score int64

	// Iterate through the sorted pairs and compute the score.
	for _, pair := range pairs {
		idx := pair.Index

		// Skip if the current element is already marked.
		if marked[idx] {
			continue
		}

		// Add the value to the score and mark the current index.
		score += int64(pair.Value)
		marked[idx] = true

		// Mark adjacent indices if they exist.
		if idx > 0 {
			marked[idx-1] = true
		}
		if idx < len(nums)-1 {
			marked[idx+1] = true
		}
	}

	return score
}
