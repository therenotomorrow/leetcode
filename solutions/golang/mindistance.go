package golang

import (
	"math"
)

func minDistance(_ int, _ int, tree []int, squirrel []int, nuts [][]int) int {
	const double = 2

	total := 0
	first := math.MaxInt

	for _, nut := range nuts {
		total += double * Manhattan(nut[0], nut[1], tree[0], tree[1])
		nutVsSquirrel := Manhattan(nut[0], nut[1], squirrel[0], squirrel[1])
		nutVsTree := Manhattan(nut[0], nut[1], tree[0], tree[1])
		first = Min(first, nutVsSquirrel-nutVsTree)
	}

	return total + first
}
