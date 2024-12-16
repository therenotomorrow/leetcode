package golang

import (
	"slices"
)

func getFinalState(nums []int, k int, multiplier int) []int {
	for ; k > 0; k-- {
		idx := slices.Index(nums, Min(nums...))

		nums[idx] *= multiplier
	}

	return nums
}
