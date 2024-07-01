package golang

func isPalindrome1(num int) bool {
	if num < 0 {
		// negative values cannot be a palindrome
		return false
	}

	res := 0
	tmp := num

	for tmp > 0 {
		res *= Digits
		res += tmp % Digits
		tmp /= Digits
	}

	return num == res
}
