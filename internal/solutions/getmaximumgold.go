package solutions

import "github.com/therenotomorrow/leetcode/pkg/mathfunc"

func getMaximumGold(grid [][]int) int {
	var (
		gold      = 0
		visited   = make([][]bool, len(grid))
		backtrack func(grid [][]int, currI int, currJ int) int
	)

	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	backtrack = func(grid [][]int, currI int, currJ int) int {
		switch m := len(grid); {
		case currI < 0:
			currI = 0
		case currI >= m:
			currI = m - 1
		}

		switch n := len(grid[0]); {
		case currJ < 0:
			currJ = 0
		case currJ >= n:
			currJ = n - 1
		}

		if grid[currI][currJ] == 0 || visited[currI][currJ] {
			return 0
		}

		visited[currI][currJ] = true

		up := backtrack(grid, currI-1, currJ)
		down := backtrack(grid, currI+1, currJ)
		left := backtrack(grid, currI, currJ-1)
		right := backtrack(grid, currI, currJ+1)

		visited[currI][currJ] = false

		return mathfunc.Max(left, right, up, down) + grid[currI][currJ]
	}

	for i, row := range grid {
		for j := range row {
			gold = mathfunc.Max(backtrack(grid, i, j), gold)
		}
	}

	return gold
}
