package golang

func findMatrix(nums []int) [][]int {
	cnt := [201]int{}

	for _, num := range nums {
		cnt[num]++
	}

	matrix := make([][]int, 0)

	for num, times := range cnt {
		for i := 0; i < times; i++ {
			if len(matrix) < i+1 {
				matrix = append(matrix, make([]int, 0))
			}

			matrix[i] = append(matrix[i], num)
		}
	}

	return matrix
}
