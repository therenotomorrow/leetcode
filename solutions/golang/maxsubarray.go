package golang

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := nums[0]

	for _, num := range nums[1:] {
		// if num > currSum then we will use 1 element array
		currSum = Max(num, currSum+num)
		maxSum = Max(maxSum, currSum)
	}

	return maxSum
}
