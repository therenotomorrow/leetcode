package golang

func minSubArrayLen(target int, nums []int) int {
	i := 0
	curSum := 0
	minLen := len(nums) + 1

	for j := range len(nums) {
		curSum += nums[j]

		for ; curSum >= target; i++ {
			curSum -= nums[i]
			minLen = Min(minLen, j-i+1)
		}
	}

	if minLen > len(nums) {
		return 0
	}

	return minLen
}
