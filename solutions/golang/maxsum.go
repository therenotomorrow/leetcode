package golang

import (
	"maps"
	"slices"
)

func maxSum(nums []int) int {
	ans := nums[0]
	cnt := make(map[int]struct{})

	for _, num := range nums {
		if num > 0 {
			cnt[num] = struct{}{}
		}

		ans = Max(ans, num)
	}

	if len(cnt) == 0 {
		return ans // because all values are negative
	}

	return Sum(slices.Collect(maps.Keys(cnt))...)
}
