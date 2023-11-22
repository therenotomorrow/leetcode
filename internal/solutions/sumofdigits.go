package solutions

import "github.com/therenotomorrow/leetcode/pkg/mathfunc"

func sumOfDigits(nums []int) int {
	minNum := mathfunc.Min(nums...)

	sum := 0

	for minNum > 0 {
		sum += minNum % 10
		minNum /= 10
	}

	return sum%2 ^ 1
}
