package golang

func shareCandies(candies []int, k int) int {
	mine := make(map[int]int)

	for _, candy := range candies {
		mine[candy]++
	}

	for _, candy := range candies[:k] {
		mine[candy]--

		if mine[candy] == 0 {
			delete(mine, candy)
		}
	}

	ans := len(mine)

	for i, j := 0, k; j < len(candies); i, j = i+1, j+1 {
		del := candies[j]
		add := candies[i]

		mine[add]++
		mine[del]--

		if mine[del] == 0 {
			delete(mine, del)
		}

		ans = Max(ans, len(mine))
	}

	return ans
}
