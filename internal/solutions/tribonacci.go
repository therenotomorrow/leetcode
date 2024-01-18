package solutions

func tribonacci(n int) int {
	arr := [3]int{0, 1, 1}

	for ; n > 2; n-- {
		arr[2], arr[1], arr[0] = arr[2]+arr[1]+arr[0], arr[2], arr[1]
	}

	return arr[n]
}
