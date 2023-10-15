package spiralOrder

func spiralOrder(matrix [][]int) []int {
	var (
		i, j, k int
		rows    = len(matrix)
		cols    = len(matrix[0])
		ans     = make([]int, rows*cols)
		offset  = 1
	)

	for j = -1; cols*rows > 0; {
		for dx := 0; dx < cols; dx++ {
			j += offset
			ans[k] = matrix[i][j]
			k++
		}

		rows--

		for dy := 0; dy < rows; dy++ {
			i += offset
			ans[k] = matrix[i][j]
			k++
		}

		cols--

		offset = -offset
	}

	return ans
}
