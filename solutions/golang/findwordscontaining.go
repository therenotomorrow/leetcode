package golang

func findWordsContaining(words []string, x byte) []int {
	ans := make([]int, 0)
	xxx := rune(x)

	for i, word := range words {
		for _, char := range word {
			if char == xxx {
				ans = append(ans, i)

				break
			}
		}
	}

	return ans
}
