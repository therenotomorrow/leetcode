package golang

func strangePrinter(text string) int {
	var (
		distinct = make([]rune, 0)
		memo     [][]int // alternative of `cache.NewCache()`
		dynamic  func(i int, j int) int
	)

	dynamic = func(i int, j int) int {
		if i > j {
			return 0
		}

		if val := memo[i][j]; val != -1 {
			return val
		}

		minActions := 1 + dynamic(i+1, j)

		for k := i + 1; k < j+1; k++ {
			if distinct[i] != distinct[k] {
				continue
			}

			minActions = Min(minActions, dynamic(i, k-1)+dynamic(k+1, j))
		}

		memo[i][j] = minActions

		return minActions
	}

	prev := rune(0)
	for _, curr := range text {
		if prev != curr {
			distinct = append(distinct, curr)
			prev = curr
		}
	}

	memo = make([][]int, len(distinct)+1)
	for i := range memo {
		memo[i] = make([]int, len(distinct)+1)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return dynamic(0, len(distinct)-1)
}
