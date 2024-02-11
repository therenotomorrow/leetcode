package solutions

import (
	"github.com/therenotomorrow/leetcode/pkg/cache"
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func cherryPickup(grid [][]int) int {
	var (
		c       = cache.NewCache()
		dynamic func(i int, robot1 int, robot2 int) int
	)

	dynamic = func(i int, robot1 int, robot2 int) int {
		if i >= len(grid) {
			return 0
		}

		if robot1 < 0 || robot1 >= len(grid[0]) || robot2 < 0 || robot2 >= len(grid[0]) {
			return -1
		}

		if val, ok := c.Load(i, robot1, robot2); ok {
			return val
		}

		cherries := grid[i][robot1]
		// both robots in the same cell, only one takes
		if robot1 != robot2 {
			cherries += grid[i][robot2]
		}

		nextMove := -1

		for r1 := robot1 - 1; r1 <= robot1+1; r1++ {
			for r2 := robot2 - 1; r2 <= robot2+1; r2++ {
				nextMove = mathfunc.Max(nextMove, dynamic(i+1, r1, r2))
			}
		}

		cherries += nextMove

		c.Save(cherries, i, robot1, robot2)

		return cherries
	}

	return dynamic(0, 0, len(grid[0])-1)
}
