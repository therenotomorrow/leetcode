package golang

func isReachableAtTime(sx int, sy int, fx int, fy int, tpl int) bool {
	xLen := Abs(fx - sx)
	yLen := Abs(fy - sy)

	// we are at the same point
	if xLen == 0 && yLen == 0 {
		if tpl == 1 {
			return false
		}
	}

	return Max(xLen, yLen) <= tpl
}
