package minScore

import (
	"sort"
)

func minScore(grid [][]int) [][]int {
	rows := len(grid)
	cols := len(grid[0])
	nums := make([][]int, 0)
	rowMax := make([]int, rows)
	colMax := make([]int, cols)

	for i := range rowMax {
		rowMax[i] = 1
	}

	for i := range colMax {
		colMax[i] = 1
	}

	for i := range rows {
		for j := range cols {
			nums = append(nums, []int{grid[i][j], i, j})
		}
	}

	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	for _, num := range nums {
		x := num[1]
		y := num[2]

		newValue := max(rowMax[x], colMax[y])
		grid[x][y] = newValue

		rowMax[x] = newValue + 1
		colMax[y] = newValue + 1
	}

	return grid
}
