package golang

import "strconv"

func sequentialDigits(low int, high int) []int {
	pattern := "123456789"
	numbers := make([]int, 0)

	// Digits - max length of the number by task description
	for r := 2; r < Digits; r++ {
		for l := range Digits - r {
			num, _ := strconv.Atoi(pattern[l : l+r])

			if low <= num && num <= high {
				numbers = append(numbers, num)
			}
		}
	}

	return numbers
}
