package golang

import (
	"math"
)

func countSymmetricIntegers(low int, high int) int {
	cnt := 0
	curr := low

	sumDigits := func(num int) int {
		sum := 0

		for num > 0 {
			sum += num % Digits
			num /= Digits
		}

		return sum
	}

	countDigits := func(num int) int {
		cnt := 0

		for num > 0 {
			num /= Digits
			cnt++
		}

		return cnt
	}

	for ; curr <= high; curr++ {
		digits := countDigits(curr)
		if digits%2 != 0 {
			continue
		}

		pow := int(math.Pow10(digits / Half))

		if sumDigits(curr/pow) == sumDigits(curr%pow) {
			cnt++
		}
	}

	return cnt
}
