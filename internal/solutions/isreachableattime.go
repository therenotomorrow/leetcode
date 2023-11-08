package solutions

import "github.com/therenotomorrow/leetcode/pkg/mathfunc"

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	x := mathfunc.Abs(fx - sx)
	y := mathfunc.Abs(fy - sy)

	// we are at the same point
	if x == 0 && y == 0 {
		if t == 1 {
			return false
		}
	}

	return mathfunc.Max(x, y) <= t
}
