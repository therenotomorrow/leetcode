package golang

func validPalindrome(str string) bool {
	left := 0
	right := len(str) - 1

	checker := func(left int, right int) bool {
		for left < right {
			if str[left] != str[right] {
				return false
			}

			left++
			right--
		}

		return true
	}

	for left < right {
		if str[left] != str[right] {
			return checker(left+1, right) || checker(left, right-1)
		}

		left++
		right--
	}

	return true
}
