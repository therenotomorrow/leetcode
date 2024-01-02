package isSubsequence

func isSubsequence(s string, t string) bool {
	curr := 0

	for i := 0; i < len(t) && curr < len(s); i++ {
		if t[i] == s[curr] {
			curr++
		}
	}

	return curr == len(s)
}
