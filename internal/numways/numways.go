package numways

import "github.com/therenotomorrow/leetcode/pkg/cache"

const MOD int = 1e9 + 7

func numWays(steps int, arrLen int) int {
	var (
		c       = cache.New()
		dynamic func(currPos int, stepsRemain int) (waysCnt int)
	)

	dynamic = func(currPos int, stepsRemain int) (waysCnt int) {
		if stepsRemain == 0 {
			if currPos == 0 {
				waysCnt = 1
			}

			return waysCnt
		}

		if val, ok := c.Load(currPos, stepsRemain); ok {
			return val
		}

		waysCnt = dynamic(currPos, stepsRemain-1) // don't move

		if currPos > 0 {
			waysCnt = (waysCnt + dynamic(currPos-1, stepsRemain-1)) % MOD
		}

		if currPos < arrLen-1 {
			waysCnt = (waysCnt + dynamic(currPos+1, stepsRemain-1)) % MOD
		}

		return c.Save(waysCnt, currPos, stepsRemain)
	}

	return dynamic(0, steps)
}
