package minDays

func minDays(grid [][]int) int {
	counter := func() int {
		var (
			cnt      = 0
			dirs     = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
			visited  = make([][]bool, len(grid))
			explorer func(row int, col int)
		)

		explorer = func(row int, col int) {
			visited[row][col] = true

			for _, direction := range dirs {
				i := row + direction[0]
				j := col + direction[1]

				if i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0]) && grid[i][j] == 1 && !visited[i][j] {
					explorer(i, j)
				}
			}
		}

		for i := range visited {
			visited[i] = make([]bool, len(grid[0]))
		}

		for row := range len(grid) {
			for col := range len(grid[0]) {
				if !visited[row][col] && grid[row][col] == 1 {
					explorer(row, col)
					cnt++
				}
			}
		}

		return cnt
	}

	if counter() != 1 {
		return 0
	}

	for row := range len(grid) {
		for col := range len(grid[0]) {
			if grid[row][col] == 0 {
				continue
			}

			grid[row][col] = 0

			if counter() != 1 {
				return 1
			}

			grid[row][col] = 1
		}
	}

	return 2
}
