package golang

import "sort"

func reductionOperations(nums []int) int {
	sort.Ints(nums)

	var ans, inc int

	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			inc++
		}

		ans += inc
	}

	return ans
}
