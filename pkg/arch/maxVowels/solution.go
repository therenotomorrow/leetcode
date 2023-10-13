package maxVowels

func isVowel(b byte) int {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return 1
	default:
		return 0
	}
}

func maxVowels(s string, k int) int {
	res := 0
	for i := range s[:k] {
		res += isVowel(s[i])
	}

	max := res
	for i, j := 0, k; j < len(s); {
		max -= isVowel(s[i])
		max += isVowel(s[j])

		if max > res {
			res = max
		}

		i++
		j++
	}

	return res
}
