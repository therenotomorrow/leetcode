package golang

import "unicode"

func isValid(word string) bool {
	vowels := NewSet[rune]()
	vowels.Populate('a', 'e', 'i', 'o', 'u')

	var size, vowelCnt, consonantCnt int

	for _, runa := range word {
		size++

		runa = unicode.ToLower(runa)

		if unicode.IsDigit(runa) {
			continue
		}

		if !unicode.IsLetter(runa) {
			return false
		}

		if vowels.Contains(runa) {
			vowelCnt++
		} else {
			consonantCnt++
		}
	}

	return size >= 3 && vowelCnt > 0 && consonantCnt > 0
}
