package solutions

import (
	"math"

	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func assignBikes2(workers [][]int, bikes [][]int) int {
	var (
		dynamic func(worker int, currSum int) int
		taken   = make([]bool, len(bikes))
		minSum  = math.MaxInt
	)

	dynamic = func(worker int, currSum int) int {
		if worker >= len(workers) {
			minSum = mathfunc.Min(minSum, currSum)

			return minSum
		}

		if currSum >= minSum {
			return minSum
		}

		for bike := range bikes {
			if taken[bike] {
				continue
			}

			taken[bike] = true
			currSum += mathfunc.Manhattan(workers[worker][0], workers[worker][1], bikes[bike][0], bikes[bike][1])
			minSum = dynamic(worker+1, currSum)
			taken[bike] = false
		}

		return minSum
	}

	return dynamic(0, 0)
}
