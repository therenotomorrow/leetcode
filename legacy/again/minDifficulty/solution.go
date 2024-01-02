package minDifficulty

import (
	"math"

	"github.com/therenotomorrow/leetcode/pkg/cache"
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func minDifficulty(jobDifficulty []int, d int) int {
	var (
		c       = cache.NewCache()
		dynamic func(curr int, remain int) (diff int)
	)

	if len(jobDifficulty) < d {
		return -1
	}

	dynamic = func(curr int, remain int) (diff int) {
		if val, ok := c.Load(curr, remain); ok {
			return val
		}

		if remain == 1 {
			return mathfunc.Max(jobDifficulty[curr:]...)
		}

		diff = math.MaxInt
		dailyMaxDiff := 0

		for job := curr; job < len(jobDifficulty)-remain+1; job++ {
			dailyMaxDiff = mathfunc.Max(dailyMaxDiff, jobDifficulty[job])
			diff = mathfunc.Min(diff, dailyMaxDiff+dynamic(job+1, remain-1))
		}

		c.Save(diff, curr, remain)

		return diff
	}

	return dynamic(0, d)
}
