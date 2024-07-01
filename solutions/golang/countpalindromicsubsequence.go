package golang

import (
	"strings"
)

func countPalindromicSubsequence(str string) int {
	letters := NewSet[rune]()

	for _, r := range str {
		letters.Add(r)
	}

	cnt := 0

	for _, let := range letters.Values() {
		left := strings.Index(str, string(let))
		right := strings.LastIndex(str, string(let))
		between := NewSet[rune]()

		for k := left + 1; k < right; k++ {
			between.Add(rune(str[k]))
		}

		cnt += between.Size()
	}

	return cnt
}
