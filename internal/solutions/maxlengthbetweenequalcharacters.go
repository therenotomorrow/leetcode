package solutions

import "github.com/therenotomorrow/leetcode/pkg/mathfunc"

func maxLengthBetweenEqualCharacters(s string) int {
	indexes := make(map[rune]int)
	maxLen := -1

	for i, r := range s {
		j, ok := indexes[r]

		if ok {
			maxLen = mathfunc.Max(maxLen, i-j-1)
		} else {
			indexes[r] = i
		}
	}

	return maxLen
}
