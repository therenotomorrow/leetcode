package golang

func lengthOfLongestSubstring(s string) int {
	var (
		maxLen, left int
		seen         = make(map[rune]int)
	)

	for right, r := range s {
		left = Max(left, seen[r])
		maxLen = Max(maxLen, right-left+1) // non-repeating window
		seen[r] = right + 1
	}

	return maxLen
}
