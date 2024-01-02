package solutions

import (
	"math"

	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func minDistance(_ int, _ int, tree []int, squirrel []int, nuts [][]int) int {
	total := 0
	first := math.MaxInt

	for _, nut := range nuts {
		total += 2 * mathfunc.Manhattan(nut[0], nut[1], tree[0], tree[1])
		nutVsSquirrel := mathfunc.Manhattan(nut[0], nut[1], squirrel[0], squirrel[1])
		nutVsTree := mathfunc.Manhattan(nut[0], nut[1], tree[0], tree[1])
		first = mathfunc.Min(first, nutVsSquirrel-nutVsTree)
	}

	return total + first
}
