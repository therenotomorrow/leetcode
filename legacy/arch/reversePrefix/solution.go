package reversePrefix

func reversePrefix(word string, ch byte) string {
	res := make([]byte, len(word))
	idx := 0

	for i := range word {
		if word[i] == ch {
			idx = i
			break
		}
	}

	for i := idx; i >= 0; i-- {
		res[idx-i] = word[i]
	}

	for i := idx + 1; i < len(word); i++ {
		res[i] = word[i]
	}

	return string(res)
}
