package findMaxConsecutiveOnes

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxConsecutiveOnes(nums []int) int {
	var ans, tmp int

	for _, num := range nums {
		if num == 0 {
			tmp = 0
		}

		tmp += num
		ans = max(ans, tmp)
	}

	return ans
}
