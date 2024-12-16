package maxLength

func maxLength(ribbons []int, k int) int {
	// Helper function to check if we can cut at least k ribbons of length x
	canCut := func(x int) bool {
		count := 0
		for _, length := range ribbons {
			count += length / x
			if count >= k {
				return true
			}
		}
		return false
	}

	// Binary search for the maximum ribbon length
	left, right := 1, 0
	for _, length := range ribbons {
		if length > right {
			right = length
		}
	}

	maxLen := 0
	for left <= right {
		mid := left + (right-left)/2
		if canCut(mid) {
			maxLen = mid
			left = mid + 1 // Try for a longer length
		} else {
			right = mid - 1 // Try for a shorter length
		}
	}

	return maxLen
}
