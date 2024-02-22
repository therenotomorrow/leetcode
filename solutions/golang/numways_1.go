package golang

func numWays1(steps int, arrLen int) int {
	var (
		c       = NewCache()
		dynamic func(currPos int, stepsRemain int) int
	)

	dynamic = func(currPos int, stepsRemain int) int {
		if stepsRemain == 0 {
			if currPos == 0 {
				return 1
			}

			return 0
		}

		if val, ok := c.Load(currPos, stepsRemain); ok {
			return val
		}

		waysCnt := dynamic(currPos, stepsRemain-1) // don't move

		if currPos > 0 {
			waysCnt = (waysCnt + dynamic(currPos-1, stepsRemain-1)) % MOD
		}

		if currPos < arrLen-1 {
			waysCnt = (waysCnt + dynamic(currPos+1, stepsRemain-1)) % MOD
		}

		c.Save(waysCnt, currPos, stepsRemain)

		return waysCnt
	}

	return dynamic(0, steps)
}
