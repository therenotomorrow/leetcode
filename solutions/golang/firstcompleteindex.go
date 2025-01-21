package golang

func firstCompleteIndex(arr []int, mat [][]int) int {
	rows := len(mat)
	cols := len(mat[0])

	rowsCnt := make([]int, rows)
	colsCnt := make([]int, cols)

	positions := make(map[int][]int)

	for i, row := range mat {
		for j, num := range row {
			positions[num] = []int{i, j}
		}
	}

	for i, num := range arr {
		coord := positions[num]
		row, col := coord[0], coord[1]

		rowsCnt[row]++
		colsCnt[col]++

		if rowsCnt[row] == cols || colsCnt[col] == rows {
			return i
		}
	}

	return -1
}
