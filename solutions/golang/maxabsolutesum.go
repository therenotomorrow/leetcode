package golang

func maxAbsoluteSum(nums []int) int {
	minPrefix := 0
	maxPrefix := 0
	prefixSum := 0

	for _, num := range nums {
		prefixSum += num

		minPrefix = Min(minPrefix, prefixSum)
		maxPrefix = Max(maxPrefix, prefixSum)
	}

	return maxPrefix - minPrefix
}
