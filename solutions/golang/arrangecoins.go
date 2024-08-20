package golang

func arrangeCoins(n int) int {
	ans := 0

	for step := 1; n >= 0; step++ {
		n -= step
		ans++
	}

	return ans - 1
}
