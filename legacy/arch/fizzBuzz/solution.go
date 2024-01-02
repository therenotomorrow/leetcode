package fizzBuzz

import "strconv"

func fizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 0; i < n; i++ {
		var val string

		switch j := i + 1; {
		case j%15 == 0:
			val = "FizzBuzz"
		case j%3 == 0:
			val = "Fizz"
		case j%5 == 0:
			val = "Buzz"
		default:
			val = strconv.Itoa(j)
		}

		result[i] = val
	}

	return result
}
