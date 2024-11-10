package minimumSubarrayLength

import (
	"math"
)

func minimumSubarrayLength(nums []int, k int) int {
	minLength := math.MaxInt
	windowStart := 0
	windowEnd := 0
	bitCounts := make([]int, 32)

	for windowEnd < len(nums) {
		updateBitCounts(bitCounts, nums[windowEnd], 1)

		for windowStart <= windowEnd && convertBitCountsToNumber(bitCounts) >= k {
			minLength = min(minLength, windowEnd-windowStart+1)

			updateBitCounts(bitCounts, nums[windowStart], -1)
			windowStart++
		}

		windowEnd++
	}

	if minLength == math.MaxInt {
		return -1
	}

	return minLength
}

func updateBitCounts(bitCounts []int, number int, delta int) {
	for bitPosition := 0; bitPosition < 32; bitPosition++ {
		if ((number >> bitPosition) & 1) != 0 {
			bitCounts[bitPosition] += delta
		}
	}
}

func convertBitCountsToNumber(bitCounts []int) int {
	result := 0
	for bitPosition := 0; bitPosition < 32; bitPosition++ {
		if bitCounts[bitPosition] != 0 {
			result |= 1 << bitPosition
		}
	}
	return result
}
