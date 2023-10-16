package solutions

import "github.com/therenotomorrow/leetcode/pkg/mathfunc"

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	left := 0
	seen := make(map[rune]int)

	for right, r := range s {
		left = mathfunc.Max(left, seen[r])
		maxLen = mathfunc.Max(maxLen, right-left+1) // non-repeating window
		seen[r] = right + 1
	}

	return maxLen
}
