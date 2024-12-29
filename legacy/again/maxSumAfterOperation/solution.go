package maxSumAfterOperation

func maxSumAfterOperation(nums []int) int {
	var (
		with    = nums[0] * nums[0]
		without = nums[0]
		maxSum  = with
	)

	for _, num := range nums[1:] {
		with = max(num*num, without+num*num, with+num)
		without = max(num, without+num)
		maxSum = max(maxSum, with)
	}

	return maxSum
}
