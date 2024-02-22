package golang

func transpose(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])

	trans := make([][]int, n)
	for i := range trans {
		trans[i] = make([]int, m)
	}

	for i, row := range matrix {
		for j, cell := range row {
			trans[j][i] = cell
		}
	}

	return trans
}
