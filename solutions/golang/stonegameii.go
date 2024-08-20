package golang

func stoneGameII(piles []int) int {
	var (
		memo    = make([][]int, len(piles)+1) // alternative of `cache.NewCache()`
		dynamic func(int, int) int
	)

	for i := range memo {
		memo[i] = make([]int, len(piles)+1)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	suffix := make([]int, len(piles)+1)
	for i := len(piles) - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}

	dynamic = func(i int, m int) int {
		if i >= len(piles) {
			return 0
		}

		if memo[i][m] != -1 {
			return memo[i][m]
		}

		currSum := 0

		for x := 1; x <= 2*m; x++ {
			currSum = Max(currSum, suffix[i]-dynamic(i+x, Max(x, m)))
		}

		memo[i][m] = currSum

		return currSum
	}

	return dynamic(0, 1)
}
