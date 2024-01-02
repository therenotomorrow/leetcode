package longestPalindrome

func oddPalindrome(s string, mid int) (left int, right int) {
	left, right = mid, mid

	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			left++
			break
		}

		left--
		right++
	}

	if left < 0 {
		left = 0
	}

	if right >= len(s) {
		right = len(s)
	}

	if s[left] != s[right-1] {
		left++
	}

	if mid-left != right-mid-1 {
		left++
	}

	return left, right
}

func evenPalindrome(s string, mid int) (left int, right int) {
	left, right = mid, mid+1

	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			left++
			break
		}

		left--
		right++
	}

	if left < 0 {
		left = 0
	}

	if right >= len(s) {
		right = len(s)
	}

	if s[left] != s[right-1] {
		left++
	}

	return left, right
}

func longestPalindrome(s string) string {
	maxLen := 0
	ans := ""

	for k := 0; k < len(s); k++ {
		i, j := oddPalindrome(s, k)
		if j-i > maxLen {
			maxLen = j - i
			ans = s[i:j]
		}

		i, j = evenPalindrome(s, k)

		if j-i > maxLen {
			maxLen = j - i
			ans = s[i:j]
		}
	}

	return ans
}
