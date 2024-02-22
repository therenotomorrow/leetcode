package golang

func firstPalindrome(words []string) string {
	for _, word := range words {
		left := 0
		right := len(word) - 1

		for left < right {
			if word[left] != word[right] {
				break
			}

			left++
			right--
		}

		if left >= right {
			return word
		}
	}

	return ""
}
