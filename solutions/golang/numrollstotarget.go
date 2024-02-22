package golang

func numRollsToTarget(n int, k int, target int) int {
	var (
		c       = NewCache()
		dynamic func(dice int, currSum int) int
	)

	dynamic = func(dice int, currSum int) int {
		if dice == n {
			if currSum == target {
				return 1
			}

			return 0
		}

		if val, ok := c.Load(dice, currSum); ok {
			return val
		}

		rolls := Min(k, target-currSum)
		rollsCnt := 0

		for i := 1; i <= rolls; i++ {
			rollsCnt += dynamic(dice+1, currSum+i)
			rollsCnt %= MOD
		}

		c.Save(rollsCnt, dice, currSum)

		return rollsCnt
	}

	return dynamic(0, 0)
}
