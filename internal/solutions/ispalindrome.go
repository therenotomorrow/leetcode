package solutions

func isPalindrome(x int) bool {
	if x < 0 {
		// negative values cannot be a palindrome
		return false
	}

	y := 0
	t := x

	for t > 0 {
		y *= 10
		y += t % 10
		t /= 10
	}

	return x == y
}
