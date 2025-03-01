package golang

func minSwaps3(data []int) int {
	var ones, ans, left int

	totalOnes := Sum(data...)

	for right, num := range data {
		ones += num
		right++

		if right-left > totalOnes {
			ones -= data[left]
			left++
		}

		ans = Max(ans, ones)
	}

	return totalOnes - ans
}
