package golang

import "math"

func gridGame(grid [][]int) int64 {
	firstRaw := Sum(grid[0]...)
	secondRow := 0
	ans := math.MaxInt

	for i, cost := range grid[0] {
		firstRaw -= cost
		ans = Min(ans, Max(firstRaw, secondRow))
		secondRow += grid[1][i]
	}

	return int64(ans)
}
