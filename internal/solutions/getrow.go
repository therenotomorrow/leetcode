package solutions

func getRow(rowIndex int) []int {
	// formula: C(n, k) = C(n, k-1) * (n - k + 1) / k

	row := make([]int, rowIndex+1)

	row[0] = 1
	for k := 1; k <= rowIndex; k++ {
		row[k] = row[k-1] * (rowIndex - k + 1) / k
	}

	return row
}
