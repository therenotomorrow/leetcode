package golang

func findMaxAverage(nums []int, k int) float64 {
	maxSum := 0
	maxSum += Sum(nums[:k]...)
	curSum := maxSum

	for i := k; i < len(nums); i++ {
		curSum += nums[i] - nums[i-k]
		maxSum = Max(maxSum, curSum)
	}

	return float64(maxSum) / float64(k)
}
