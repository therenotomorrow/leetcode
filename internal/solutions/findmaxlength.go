package solutions

import "github.com/therenotomorrow/leetcode/pkg/mathfunc"

func findMaxLength(nums []int) int {
	ans := 0
	sum := 0
	count := make(map[int]int)

	for i, num := range nums {
		if num == 1 {
			sum++
		} else {
			sum--
		}

		if sum == 0 {
			ans = mathfunc.Max(ans, i+1)
		}

		if val, ok := count[sum]; ok {
			ans = mathfunc.Max(ans, i-val)
		} else {
			count[sum] = i
		}
	}

	return ans
}
