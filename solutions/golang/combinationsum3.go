package golang

import (
	"slices"
)

func combinationSum3(k int, n int) [][]int {
	var (
		ans       = make([][]int, 0)
		backtrack func(curr int, sum int, nums []int)
	)

	backtrack = func(curr int, sum int, nums []int) {
		if len(nums) == k {
			if sum == n {
				ans = append(ans, slices.Clone(nums))
			}

			return
		}

		for next := curr + 1; next < Digits; next++ {
			nums = append(nums, next)

			backtrack(next, sum+next, nums)

			nums = nums[:len(nums)-1]
		}
	}

	backtrack(0, 0, []int{})

	return ans
}
