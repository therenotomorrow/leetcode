package plusOne

func plusOne(digits []int) []int {
	reminder := 1

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += reminder

		if digits[i] < 10 {
			reminder = 0
			break
		}

		digits[i] = 0
	}

	if reminder == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
