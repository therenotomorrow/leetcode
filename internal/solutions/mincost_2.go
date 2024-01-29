package solutions

import (
	"github.com/therenotomorrow/leetcode/pkg/cache"
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func minCost2(costs [][]int) int {
	var (
		c       = cache.NewCache()
		dynamic func(color int, house int) int
	)

	dynamic = func(color int, house int) int {
		if house >= len(costs) {
			return 0
		}

		if val, ok := c.Load(color, house); ok {
			return val
		}

		cost := costs[house][color]

		switch color {
		case 0:
			cost += mathfunc.Min(dynamic(1, house+1), dynamic(2, house+1))
		case 1:
			cost += mathfunc.Min(dynamic(0, house+1), dynamic(2, house+1))
		case 2:
			cost += mathfunc.Min(dynamic(0, house+1), dynamic(1, house+1))
		}

		c.Save(cost, color, house)

		return cost
	}

	return mathfunc.Min(dynamic(0, 0), dynamic(1, 0), dynamic(2, 0))
}
