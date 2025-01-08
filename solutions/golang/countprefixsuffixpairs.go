package golang

import "strings"

func countPrefixSuffixPairs(words []string) int {
	cnt := 0
	isPrefixAndSuffix := func(text string, prefSuf string) bool {
		return strings.HasPrefix(text, prefSuf) && strings.HasSuffix(text, prefSuf)
	}

	for i, word1 := range words {
		for j, word2 := range words {
			if i < j && isPrefixAndSuffix(word2, word1) {
				cnt++
			}
		}
	}

	return cnt
}
