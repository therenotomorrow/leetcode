package golang

func maxDistance2(s string, k int) int {
	var deltaX, deltaY, ans int

	for i, direction := range s {
		switch direction {
		case 'N':
			deltaY++
		case 'S':
			deltaY--
		case 'E':
			deltaX++
		case 'W':
			deltaX--
		}

		ans = Max(ans, Min(Abs(deltaX)+Abs(deltaY)+Double*k, i+1))
	}

	return ans
}
