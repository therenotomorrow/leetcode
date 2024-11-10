package golang

func minEnd(n int, x int) int64 {
	ans := x

	for ; n > 1; n-- {
		ans = (ans + 1) | x
	}

	return int64(ans)
}
