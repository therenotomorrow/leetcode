package shortestPalindrome

import "slices"

func shortestPalindrome(s string) string {
	length := len(s)
	if length == 0 {
		return s
	}

	left := 0
	for right := length - 1; right >= 0; right-- {
		if s[right] == s[left] {
			left++
		}
	}

	if left == length {
		return s
	}

	nonPalindromeSuffix := []rune(s[left:])
	reverseSuffix := make([]rune, len(nonPalindromeSuffix))
	copy(reverseSuffix, nonPalindromeSuffix)

	slices.Reverse(reverseSuffix)

	return string(reverseSuffix) + shortestPalindrome(s[:left]) + string(nonPalindromeSuffix)
}
