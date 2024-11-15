package golang

func getMaximumGenerated(n int) int {
	const size = 2

	ans := 1
	arr := []int{0, 1}

	if n < size {
		return arr[n]
	}

	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			arr = append(arr, arr[i/2])
		} else {
			arr = append(arr, arr[i/2]+arr[i/2+1])
		}

		ans = Max(ans, arr[i])
	}

	return ans
}
