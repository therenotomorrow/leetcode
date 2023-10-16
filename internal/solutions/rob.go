package solutions

import "github.com/therenotomorrow/leetcode/pkg/mathfunc"

func rob(nums []int) int {
	worth := make([]int, 0)
	worth = append(worth, 0, nums[0])

	for i := 1; i < len(nums); i++ {
		next := mathfunc.Max(worth[i], worth[i-1]+nums[i])
		worth = append(worth, next)
	}

	return worth[len(nums)]
}
