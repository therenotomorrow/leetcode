package golang

import (
	"sort"
)

func check(nums []int) bool {
	n := len(nums)
	nums = append(nums, nums...)

	for i := range n {
		if sort.IsSorted(sort.IntSlice(nums[i : i+n])) {
			return true
		}
	}

	return false
}
