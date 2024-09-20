package golang

import (
	"slices"
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	cnt := make(map[string]int)

	for _, word := range slices.Concat(strings.Fields(s1), strings.Fields(s2)) {
		cnt[word]++
	}

	ans := make([]string, 0)

	for word, times := range cnt {
		if times == 1 {
			ans = append(ans, word)
		}
	}

	return ans
}
