package golang

import (
	"sort"
)

func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)

	ans := 0
	left := 0

	for right, num := range nums {
		for num-nums[left] > 2*k {
			left++
		}

		ans = Max(ans, right-left+1)
	}

	return ans
}
