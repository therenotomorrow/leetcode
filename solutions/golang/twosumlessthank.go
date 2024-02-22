package golang

import (
	"sort"
)

func twoSumLessThanK(nums []int, k int) int {
	sort.Ints(nums)

	ans := -1
	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[right] + nums[left]

		if sum < k {
			ans = Max(ans, sum)
			left++
		} else {
			right--
		}
	}

	return ans
}
