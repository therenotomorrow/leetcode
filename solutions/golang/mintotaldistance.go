package golang

import "sort"

func minTotalDistance(grid [][]int) int {
	rows := make([]int, 0)
	cols := make([]int, 0)

	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	// because we collect cols in shuffle queue
	sort.Ints(cols)

	return calcDist(rows) + calcDist(cols)
}

func calcDist(arr []int) int {
	dist := 0

	for i, j := 0, len(arr)-1; i < j; {
		dist += arr[j] - arr[i]
		i++
		j--
	}

	return dist
}
