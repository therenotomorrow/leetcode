package golang

func probabilityOfHeads(prob []float64, target int) float64 {
	var (
		memo    = make([][]float64, len(prob)+1) // alternative of `cache.NewCache()`
		dynamic func(curr int, target int) float64
	)

	for i := range memo {
		memo[i] = make([]float64, target+1)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	dynamic = func(curr int, goal int) float64 {
		if goal < 0 {
			return 0
		}

		if curr == len(prob) {
			if goal == 0 {
				return 1
			}

			return 0
		}

		if val := memo[curr][goal]; val != -1 {
			return val
		}

		// first - we toss head, second - negative probability (sum equals 1)
		memo[curr][goal] = dynamic(curr+1, goal-1)*prob[curr] + dynamic(curr+1, goal)*(1-prob[curr])

		return memo[curr][goal]
	}

	return dynamic(0, target)
}
