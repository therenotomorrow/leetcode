package golang

func findLengthOfLCIS(nums []int) int {
	maxLen := 1

	for i := range nums {
		cntLen := 1
		prev := nums[i]

		for _, curr := range nums[i+1:] {
			if curr <= prev {
				break
			}

			cntLen++
			prev = curr
		}

		maxLen = Max(maxLen, cntLen)
	}

	return maxLen
}
