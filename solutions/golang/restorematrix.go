package golang

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	matrix := make([][]int, len(rowSum))
	for i := range matrix {
		matrix[i] = make([]int, len(colSum))
	}

	currRowSum := make([]int, len(rowSum))
	currColSum := make([]int, len(colSum))

	for i := range rowSum {
		for j := range colSum {
			matrix[i][j] = Min(rowSum[i]-currRowSum[i], colSum[j]-currColSum[j])

			currRowSum[i] += matrix[i][j]
			currColSum[j] += matrix[i][j]
		}
	}

	return matrix
}
