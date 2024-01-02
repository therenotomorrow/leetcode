package solutions

import "github.com/therenotomorrow/leetcode/pkg/mathfunc"

func getSumAbsoluteDifferences(nums []int) []int {
	var (
		ans               = make([]int, len(nums))
		totalSum          = mathfunc.Sum(nums...)
		leftSum, rightSum int
	)

	for i, num := range nums {
		rightSum = totalSum - leftSum - num

		leftCnt := i
		rightCnt := len(nums) - 1 - leftCnt

		leftTotal := leftCnt*num - leftSum
		rightTotal := rightSum - rightCnt*num

		ans[i] = leftTotal + rightTotal
		leftSum += num
	}

	return ans
}
