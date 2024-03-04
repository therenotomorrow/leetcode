package golang

func findMaxConsecutiveOnes2(nums []int) int {
	var ans, left, zeros int

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
