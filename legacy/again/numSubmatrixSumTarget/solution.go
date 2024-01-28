package numSubmatrixSumTarget

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	prefixSums := make([][]int, rows+1)
	prefixSums[0] = make([]int, cols+1)

	for i := 1; i < rows+1; i++ {
		prefixSums[i] = make([]int, cols+1)

		for j := 1; j < cols+1; j++ {
			prefixSums[i][j] = prefixSums[i-1][j] + prefixSums[i][j-1] - prefixSums[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	count := 0
	currSum := 0

	cnt := make(map[int]int)

	for r1 := 1; r1 < rows+1; r1++ {
		for r2 := r1; r2 < rows+1; r2++ {
			clear(cnt)
			cnt[0] = 1

			for col := 1; col < cols+1; col++ {
				currSum = prefixSums[r2][col] - prefixSums[r1-1][col]
				count += cnt[currSum-target]
				cnt[currSum] = cnt[currSum] + 1
			}
		}
	}

	return count
}
