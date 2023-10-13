package mergeAlternately

func mergeAlternately(word1 string, word2 string) string {
	l1, l2 := len(word1), len(word2)

	result := make([]byte, 0, l1+l2)
	for i := 0; i < l1 && i < l2; i++ {
		result = append(result, word1[i], word2[i])
	}

	if l1 > l2 {
		result = append(result, word1[l2:]...)
	} else {
		result = append(result, word2[l1:]...)
	}

	return string(result)
}
