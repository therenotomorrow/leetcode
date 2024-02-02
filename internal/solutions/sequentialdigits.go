package solutions

import "strconv"

func sequentialDigits(low int, high int) []int {
	pattern := "123456789"
	numbers := make([]int, 0)

	// 10 - max length of the number by task description
	for r := 2; r < 10; r++ {
		for l := 0; l < 10-r; l++ {
			num, _ := strconv.Atoi(pattern[l : l+r])

			if low <= num && num <= high {
				numbers = append(numbers, num)
			}
		}
	}

	return numbers
}
