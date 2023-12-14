package solutions

func onesMinusZeros(grid [][]int) [][]int {
	n := len(grid)
	m := len(grid[0])

	onesRows := make([]int, n)
	onesCols := make([]int, m)

	for i, row := range grid {
		for j, cell := range row {
			onesRows[i] += cell
			onesCols[j] += cell
		}
	}

	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}

	for i, row := range grid {
		for j := range row {
			ans[i][j] = onesRows[i] + onesCols[j] - (n - onesRows[i]) - (m - onesCols[j])
		}
	}

	return ans
}
