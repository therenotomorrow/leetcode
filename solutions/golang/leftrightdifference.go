package golang

func leftRightDifference(nums []int) []int {
	const two = 2

	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))

	for i, j := 1, len(nums)-two; i < len(nums) && j >= 0; {
		leftSum[i] = leftSum[i-1] + nums[i-1]
		rightSum[j] = rightSum[j+1] + nums[j+1]

		i++
		j--
	}

	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = Abs(leftSum[i] - rightSum[i])
	}

	return ans
}
