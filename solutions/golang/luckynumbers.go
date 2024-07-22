package golang

func luckyNumbers(matrix [][]int) []int {
	set := NewSet[int]()
	ans := make([]int, 0)

	for _, row := range matrix {
		set.Add(Min(row...))
	}

	for j := range len(matrix[0]) {
		tmp := make([]int, len(matrix))

		for i := range len(matrix) {
			tmp[i] = matrix[i][j]
		}

		if mx := Max(tmp...); set.Contains(mx) {
			ans = append(ans, mx)
		}
	}

	return ans
}
