package golang

func getLongestSubsequence(words []string, groups []int) []string {
	ans := make([]string, 0)

	for i, word := range words {
		if i == 0 || groups[i] != groups[i-1] {
			ans = append(ans, word)
		}
	}

	return ans
}
