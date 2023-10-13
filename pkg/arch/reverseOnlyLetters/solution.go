package reverseOnlyLetters

func isAlpha(b byte) bool {
	return 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z'
}

func reverseOnlyLetters(s string) string {
	res := make([]byte, len(s))
	l := 0
	r := len(s) - 1

	for {
		for ; l < len(s) && !isAlpha(s[l]); l++ {
			res[l] = s[l]
		}

		for ; r >= 0 && !isAlpha(s[r]); r-- {
			res[r] = s[r]
		}

		if l > r {
			break
		}

		res[l], res[r] = s[r], s[l]
		l++
		r--
	}

	return string(res)
}
