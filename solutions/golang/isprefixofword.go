package golang

import (
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	for i, word := range strings.Fields(sentence) {
		if strings.HasPrefix(word, searchWord) {
			return i + 1
		}
	}

	return -1
}
