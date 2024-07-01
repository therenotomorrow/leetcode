package golang

func reverse(num int) int {
	res := 0

	for num != 0 {
		digit := num % Digits

		if res > MaxInt32Overflow || res < MinInt32Overflow {
			return 0
		}

		num /= Digits
		res = res*Digits + digit
	}

	return res
}
