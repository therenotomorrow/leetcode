package golang

import "math"

func thirdMax(nums []int) int {
	var (
		first  = math.MinInt
		second = math.MinInt
		third  = math.MinInt
	)

	for _, num := range nums {
		if num == third || num == second || num == first {
			continue
		}

		switch {
		case num > first:
			third, second, first = second, first, num
		case num > second:
			third, second = second, num
		case num > third:
			third = num
		}
	}

	if third == math.MinInt {
		return first
	}

	return third
}
