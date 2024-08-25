package golang

import (
	"fmt"
	"strconv"
)

func fractionAddition(expression string) string { //nolint:cyclop
	var (
		numerators       = make([]int, 0)
		denominators     = make([]int, 0)
		sign             = false
		isNumerator      = true
		totalDenominator = 1
		gcd              = func(a, b int) int {
			for b != 0 {
				a, b = b, a%b
			}

			return a
		}
	)

	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case '-':
			sign = true
		case '+', '/':
			sign = false
		default:
			start := i

			// find the end of the number
			for i < len(expression) && '0' <= expression[i] && expression[i] <= '9' {
				i++
			}

			num, _ := strconv.Atoi(expression[start:i])

			if sign {
				num = -num
			}

			if isNumerator {
				numerators = append(numerators, num)
			} else {
				denominators = append(denominators, num)
				totalDenominator *= num
			}

			i-- // return `i` back because of inner cycle counter
			isNumerator = !isNumerator
		}
	}

	sum := 0
	for i, denominator := range denominators {
		sum += numerators[i] * totalDenominator / denominator
	}

	commonDenominator := Abs(gcd(sum, totalDenominator))

	sum /= commonDenominator
	totalDenominator /= commonDenominator

	return fmt.Sprintf("%d/%d", sum, totalDenominator)
}
