package golang

func getLastMoment(n int, left []int, right []int) int {
	ans := 0

	// find the max distance needed for ants moving left
	for _, ant := range left {
		ans = Max(ant, ans)
	}

	// collide don't change the time of falling
	for _, ant := range right {
		ans = Max(n-ant, ans)
	}

	return ans
}
