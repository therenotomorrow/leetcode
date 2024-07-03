package golang

import (
	"math"
	"sort"
)

func minDifference(nums []int) int {
	const size = 3

	if len(nums) <= size {
		return 0
	}

	sort.Ints(nums)

	ans := math.MaxInt
	for i := 0; i <= size; i++ {
		ans = Min(ans, nums[len(nums)-1-i]-nums[size-i])
	}

	return ans
}
