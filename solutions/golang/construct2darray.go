package golang

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != n*m {
		return [][]int{}
	}

	arr := make([][]int, m)

	for i, j := 0, 0; i < len(original); j++ {
		arr[j] = original[i : i+n]
		i += n
	}

	return arr
}
