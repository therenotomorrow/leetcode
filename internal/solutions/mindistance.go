package solutions

import (
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
	"math"
)

func nutDistance(x1 int, x2 int, y1 int, y2 int) int {
	return mathfunc.Abs(x1-x2) + mathfunc.Abs(y1-y2)
}

func minDistance(_ int, _ int, tree []int, squirrel []int, nuts [][]int) int {
	total := 0
	first := math.MaxInt

	for _, nut := range nuts {
		total += 2 * nutDistance(nut[0], tree[0], nut[1], tree[1])
		first = mathfunc.Min(first, nutDistance(nut[0], squirrel[0], nut[1], squirrel[1])-nutDistance(nut[0], tree[0], nut[1], tree[1]))
	}

	return total + first
}
