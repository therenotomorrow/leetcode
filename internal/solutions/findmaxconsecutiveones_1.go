package solutions

import "github.com/therenotomorrow/leetcode/pkg/mathfunc"

func findMaxConsecutiveOnes1(nums []int) int {
	ans := 0
	tmp := 0

	for _, num := range nums {
		if num == 0 {
			tmp = 0
		}

		tmp += num
		ans = mathfunc.Max(ans, tmp)
	}

	return ans
}
