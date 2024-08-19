package golang

func minSteps2(n int) int {
	const maxN = 1001

	var (
		cache   = NewCache()
		dynamic func(curr int, copied int) int
	)

	if n == 1 {
		return 0
	}

	dynamic = func(curr int, copied int) int {
		if curr == n {
			return 0
		}

		if curr > n {
			return maxN
		}

		if val, ok := cache.Load(curr, copied); ok {
			return val
		}

		ops := 1 + Min(dynamic(curr+copied, copied), 1+dynamic(curr+curr, curr))

		cache.Save(ops, curr, copied)

		return ops
	}

	// first operation is always copy all
	return 1 + dynamic(1, 1)
}
