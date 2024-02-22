package golang

func findMaxConsecutiveOnes2(nums []int) int {
	ans := 0
	left := 0
	zeros := 0

	for right, num := range nums {
		if num == 0 {
			zeros++
		}

		for zeros == 2 {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		ans = Max(ans, right-left+1)
	}

	return ans
}
