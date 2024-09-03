package golang

func getLucky(s string, k int) int {
	var num, digit int

	for _, runa := range s {
		digit = int(runa - 'a' + 1)

		if digit >= Digits {
			digit = digit/Digits + digit%Digits
		}

		num += digit
	}

	for ; k > 1; k-- {
		digit = 0

		for num > 0 {
			digit += num % Digits
			num /= Digits
		}

		num = digit
	}

	return num
}
