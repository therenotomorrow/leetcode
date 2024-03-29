package golang

import (
	"strings"
)

func countPalindromicSubsequence(s string) int {
	letters := NewSet[rune]()

	for _, r := range s {
		letters.Add(r)
	}

	cnt := 0

	for _, let := range letters.Values() {
		left := strings.Index(s, string(let))
		right := strings.LastIndex(s, string(let))
		between := NewSet[rune]()

		for k := left + 1; k < right; k++ {
			between.Add(rune(s[k]))
		}

		cnt += between.Size()
	}

	return cnt
}
