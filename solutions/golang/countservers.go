package golang

func countServers(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows := make([]int, len(grid[0]))
	cols := make([]int, len(grid))

	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				rows[j]++
				cols[i]++
			}
		}
	}

	cnt := 0

	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				if rows[j] > 1 || cols[i] > 1 {
					cnt++
				}
			}
		}
	}

	return cnt
}
