package golang

func findKthPositive(arr []int, k int) int {
	num := 0

	for i, j := 1, 0; k > 0; i++ {
		if j < len(arr) && arr[j] == i {
			j++
		} else {
			k--
			num = i
		}
	}

	return num
}
