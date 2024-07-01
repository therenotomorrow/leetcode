package golang

func numWays2(n int, k int) int {
	const magicTwo = 2

	var (
		cache   = NewCache()
		dynamic func(curr int, k int) int
	)

	dynamic = func(curr int, k int) int {
		switch val, ok := cache.Load(curr); {
		case ok:
			return val
		case curr == 1:
			return k
		case curr == magicTwo:
			return k * k
		}

		ans := (k - 1) * (dynamic(curr-1, k) + dynamic(curr-magicTwo, k))

		cache.Save(ans, curr)

		return ans
	}

	return dynamic(n, k)
}
