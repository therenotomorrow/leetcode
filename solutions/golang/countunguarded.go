package golang

import (
	"slices"
)

func countUnguarded(rows int, cols int, guards [][]int, walls [][]int) int { //nolint:cyclop
	const obstacle = 2

	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}

	for _, busy := range slices.Concat(guards, walls) {
		grid[busy[0]][busy[1]] = obstacle
	}

	follow := func(row int, col int) {
		for idx := col - 1; idx >= 0; idx-- {
			if grid[row][idx] > 1 {
				break
			}

			grid[row][idx] = 1
		}

		for idx := col + 1; idx < cols; idx++ {
			if grid[row][idx] > 1 {
				break
			}

			grid[row][idx] = 1
		}

		for idx := row - 1; idx >= 0; idx-- {
			if grid[idx][col] > 1 {
				break
			}

			grid[idx][col] = 1
		}

		for idx := row + 1; idx < rows; idx++ {
			if grid[idx][col] > 1 {
				break
			}

			grid[idx][col] = 1
		}
	}

	for _, guard := range guards {
		follow(guard[0], guard[1])
	}

	cnt := 0

	for _, row := range grid {
		for _, cell := range row {
			if cell == 0 {
				cnt++
			}
		}
	}

	return cnt
}
