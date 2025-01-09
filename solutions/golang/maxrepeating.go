package golang

import (
	"strings"
)

func maxRepeating(sequence string, word string) int {
	wLen := len(word)
	word = strings.Repeat(word, len(sequence)/wLen)

	for word != "" {
		if strings.Contains(sequence, word) {
			break
		}

		word = word[wLen:]
	}

	return len(word) / wLen
}
