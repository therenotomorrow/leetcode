package golang

func pivotInteger(n int) int {
	// use math sqrt here instead :)
	nums := make([]int, n)
	leftSum := 0
	totalSum := nums[n-1]

	for i := 1; i <= n; i++ {
		nums[i-1] = i
		totalSum += i
	}

	for i, num := range nums {
		totalSum -= num

		if leftSum == totalSum {
			return i + 1
		}

		leftSum += num
	}

	return -1
}
