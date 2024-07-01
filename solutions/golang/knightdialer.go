package golang

func knightDialer(n int) int {
	var (
		cache   = NewCache()
		dynamic func(remain int, currDigit int) int
	)

	graph := [][]int{
		{4, 6},
		{6, 8},
		{7, 9},
		{4, 8},
		{3, 9, 0},
		{},
		{1, 7, 0},
		{2, 6},
		{1, 3},
		{2, 4},
	}

	dynamic = func(remain int, digit int) int {
		if remain == 0 {
			return 1
		}

		if val, ok := cache.Load(remain, digit); ok {
			return val
		}

		numsCnt := 0
		for _, nextDigit := range graph[digit] {
			numsCnt = (numsCnt + dynamic(remain-1, nextDigit)) % MOD
		}

		cache.Save(numsCnt, remain, digit)

		return numsCnt
	}

	numsCnt := 0
	for digit := range Digits {
		numsCnt = (numsCnt + dynamic(n-1, digit)) % MOD
	}

	return numsCnt
}
