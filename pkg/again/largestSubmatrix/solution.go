package largestSubmatrix

import (
	"math"
	"sort"
)

func largestSubmatrix(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	a := 0

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] != 0 && row > 0 {
				matrix[row][col] += matrix[row-1][col]
			}
		}

		currRow := make([]int, len(matrix[row]))
		copy(currRow, matrix[row])

		sort.Ints(currRow)

		for i := 0; i < n; i++ {
			a = Max(a, currRow[i]*(n-i))
		}
	}

	return a
}

func Max(nums ...int) int {
	// LeetCode use Golang < 1.21
	m := math.MinInt
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}
