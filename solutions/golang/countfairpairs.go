package golang

import (
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)

	var ans [2]int

	for k, border := range []int{lower, upper + 1} {
		for i, j := 0, len(nums)-1; i < j; {
			if nums[i]+nums[j] < border {
				ans[k] += j - i
				i++
			} else {
				j--
			}
		}
	}

	return int64(ans[1] - ans[0])
}
