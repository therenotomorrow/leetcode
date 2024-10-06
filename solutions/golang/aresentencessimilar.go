package golang

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	if len(sentence2) < len(sentence1) {
		sentence1, sentence2 = sentence2, sentence1
	}

	words1 := strings.Fields(sentence1)
	words2 := strings.Fields(sentence2)
	gap := len(words2) - len(words1)

	for i, j := 0, len(words1)-1; i <= j; {
		switch {
		case words1[i] == words2[i]:
			i++
		case words1[j] == words2[gap+j]:
			j--
		default:
			return false
		}
	}

	return true
}
