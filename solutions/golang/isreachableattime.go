package golang

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	x := Abs(fx - sx)
	y := Abs(fy - sy)

	// we are at the same point
	if x == 0 && y == 0 {
		if t == 1 {
			return false
		}
	}

	return Max(x, y) <= t
}
