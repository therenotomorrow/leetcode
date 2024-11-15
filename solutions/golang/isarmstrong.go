package golang

func isArmstrong(n int) bool {
	armstrong := 0
	digits := make([]int, 0)

	for i := n; i > 0; {
		digit := i % Digits
		digits = append(digits, digit)
		i /= Digits
	}

	for _, digit := range digits {
		num := 1

		for range digits {
			num *= digit
		}

		armstrong += num
	}

	return armstrong == n
}
