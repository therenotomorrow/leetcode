package golang

func addDigits(num int) int {
	for num >= Digits {
		sum := 0

		for tmp := num; tmp > 0; tmp /= Digits {
			sum += tmp % Digits
		}

		num = sum
	}

	return num
}
