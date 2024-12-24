package maxChunksToSorted

func maxChunksToSorted(arr []int) int {
	var chunks, prefixSum, sortedPrefixSum int

	for i := range len(arr) {
		prefixSum += arr[i]
		sortedPrefixSum += i

		if prefixSum == sortedPrefixSum {
			chunks++
		}
	}

	return chunks
}
