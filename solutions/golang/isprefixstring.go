package golang

func isPrefixString(str string, words []string) bool {
	pnt := 0

	for _, word := range words {
		for j := range word {
			if pnt >= len(str) || str[pnt] != word[j] {
				return false
			}

			pnt++
		}

		// need the full word concatenation of `words`
		if pnt == len(str) {
			return true
		}
	}

	return false
}
