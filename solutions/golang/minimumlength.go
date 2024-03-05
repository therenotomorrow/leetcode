package golang

func minimumLength(s string) int {
	left := 0
	right := len(s) - 1

	for left < right && s[left] == s[right] {
		let := s[left]

		for left <= right && let == s[left] {
			left++
		}

		for right > left && let == s[right] {
			right--
		}
	}

	return right - left + 1
}
