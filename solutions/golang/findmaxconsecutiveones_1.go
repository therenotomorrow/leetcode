package golang

func findMaxConsecutiveOnes1(nums []int) int {
	var ans, tmp int

	for _, num := range nums {
		if num == 0 {
			tmp = 0
		}

		tmp += num
		ans = Max(ans, tmp)
	}

	return ans
}
