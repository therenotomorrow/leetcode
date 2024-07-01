package golang

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)

	for row := range numRows {
		triangle[row] = make([]int, row+1)
		triangle[row][0] = 1

		for k := 1; k <= row; k++ {
			// formula: C(n, k) = C(n, k-1) * (n - k + 1) / k
			triangle[row][k] = triangle[row][k-1] * (row - k + 1) / k
		}
	}

	return triangle
}
