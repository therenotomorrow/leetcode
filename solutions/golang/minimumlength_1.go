package golang

func minimumLength1(str string) int {
	left := 0
	right := len(str) - 1

	for left < right && str[left] == str[right] {
		let := str[left]

		for left <= right && let == str[left] {
			left++
		}

		for right > left && let == str[right] {
			right--
		}
	}

	return right - left + 1
}
