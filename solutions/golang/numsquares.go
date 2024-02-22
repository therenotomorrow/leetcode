package golang

import (
	"math"
)

func numSquares(n int) int {
	squares := make([]int, n+1)
	for i := 1; i <= n; i++ {
		squares[i] = i * i

		if i*i > n {
			break
		}
	}

	dynamic := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dynamic[i] = math.MaxInt
	}

	for i := 1; i <= n; i++ {
		for _, square := range squares[1:] {
			if i < square {
				break
			}

			dynamic[i] = Min(dynamic[i], dynamic[i-square]+1)
		}
	}

	return dynamic[n]
}
