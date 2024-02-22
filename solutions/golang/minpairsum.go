package golang

import (
	"math"
	"sort"
)

func minPairSum(nums []int) int {
	ans := math.MinInt

	sort.Ints(nums)

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		sum := nums[i] + nums[j]
		if sum > ans {
			ans = sum
		}
	}

	return ans
}
