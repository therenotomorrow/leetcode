package golang

func numIslands(grid [][]byte) int {
	var (
		rows = len(grid)
		cols = len(grid[0])
		dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		dfs  func(row int, col int)
	)

	dfs = func(row int, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] != '1' {
			return
		}

		grid[row][col] = '0'

		for _, dir := range dirs {
			dfs(row+dir[0], col+dir[1])
		}
	}

	cnt := 0

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '1' {
				dfs(row, col)

				cnt++
			}
		}
	}

	return cnt
}
