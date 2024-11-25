package golang

import (
	"math"
)

func maxMatrixSum(matrix [][]int) int64 {
	const double = 2

	var (
		totalSum = 0
		negative = 0
		minValue = math.MaxInt
	)

	for _, row := range matrix {
		for _, val := range row {
			abs := Abs(val)

			totalSum += abs

			if val < 0 {
				negative++
			}

			minValue = Min(minValue, abs)
		}
	}

	if negative%double != 0 {
		totalSum -= double * minValue
	}

	return int64(totalSum)
}
