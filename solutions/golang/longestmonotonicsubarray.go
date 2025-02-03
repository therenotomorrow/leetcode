package golang

func longestMonotonicSubarray(nums []int) int {
	inc := 0
	dec := 0
	curr := nums[0]
	maxLen := 1

	for idx, num := range nums {
		if num > curr {
			maxLen = Max(idx-inc+1, maxLen)
		} else {
			inc = idx
		}

		if num < curr {
			maxLen = Max(idx-dec+1, maxLen)
		} else {
			dec = idx
		}

		curr = num
	}

	return maxLen
}
