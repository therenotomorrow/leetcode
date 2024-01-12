package solutions

func isPrefixString(s string, words []string) bool {
	pnt := 0

	for _, word := range words {
		for j := range word {
			if pnt >= len(s) || s[pnt] != word[j] {
				return false
			}

			pnt++
		}

		// need the full word concatenation of `words`
		if pnt == len(s) {
			return true
		}
	}

	return false
}
