package golang

import "sort"

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)

	ans := 1

	less := nums[0]
	for _, num := range nums {
		if num-less > k {
			less = num
			ans++
		}
	}

	return ans
}
