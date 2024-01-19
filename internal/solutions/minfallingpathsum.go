package solutions

import (
	"math"

	"github.com/therenotomorrow/leetcode/pkg/cache"
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func minFallingPathSum(matrix [][]int) int {
	var (
		c       = cache.NewCache()
		dynamic func(curr int, step int) int
		minSum  = math.MaxInt
	)

	dynamic = func(curr int, step int) int {
		if curr >= len(matrix) {
			return 0
		}

		if step < 0 || step >= len(matrix) {
			return math.MaxInt
		}

		if val, ok := c.Load(curr, step); ok {
			return val
		}

		sum := matrix[curr][step] + mathfunc.Min(dynamic(curr+1, step-1), dynamic(curr+1, step), dynamic(curr+1, step+1))

		c.Save(sum, curr, step)

		return sum
	}

	for i := range matrix {
		minSum = mathfunc.Min(minSum, dynamic(0, i))
	}

	return minSum
}
