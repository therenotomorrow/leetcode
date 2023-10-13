package reverseVowels

func isVowel(sym byte) bool {
	switch sym {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	case 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func reverseVowels(s string) string {
	result := make([]byte, len(s))

	for i, j := 0, len(s)-1; i <= len(s)-1; i++ {
		if !isVowel(s[i]) {
			result[i] = s[i]
			continue
		}

		if i > j {
			continue
		}

		for ; !isVowel(s[j]); j-- {
		}

		result[i], result[j] = s[j], s[i]
		j--
	}

	return string(result)
}
