package solutions

func findDiagonalOrder(nums [][]int) []int {
	diag := make(map[int][]int)
	size := 0

	for i := len(nums) - 1; i >= 0; i-- {
		for j, ceil := range nums[i] {
			diagNum := i + j
			size++

			diag[diagNum] = append(diag[diagNum], ceil)
		}
	}

	travers := make([]int, 0, size)

	for curr := 0; curr < len(diag); curr++ {
		travers = append(travers, diag[curr]...)
	}

	return travers
}
