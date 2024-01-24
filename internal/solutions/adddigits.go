package solutions

func addDigits(num int) int {
	for num >= 10 {
		sum := 0

		for tmp := num; tmp > 0; tmp /= 10 {
			sum += tmp % 10
		}

		num = sum
	}

	return num
}
