package golang

func validPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	checker := func(left int, right int) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}

			left++
			right--
		}

		return true
	}

	for left < right {
		if s[left] != s[right] {
			return checker(left+1, right) || checker(left, right-1)
		}

		left++
		right--
	}

	return true
}
