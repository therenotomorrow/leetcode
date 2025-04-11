package golang

import (
	"math"
)

func maxMatrixSum(matrix [][]int) int64 {
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

	if negative%Double != 0 {
		totalSum -= Double * minValue
	}

	return int64(totalSum)
}
