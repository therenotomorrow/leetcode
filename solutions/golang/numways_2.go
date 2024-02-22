package golang

func numWays2(n int, k int) int {
	var (
		c       = NewCache()
		dynamic func(curr int, k int) int
	)

	dynamic = func(curr int, k int) int {
		switch val, ok := c.Load(curr); {
		case ok:
			return val
		case curr == 1:
			return k
		case curr == 2:
			return k * k
		}

		ans := (k - 1) * (dynamic(curr-1, k) + dynamic(curr-2, k))

		c.Save(ans, curr)

		return ans
	}

	return dynamic(n, k)
}
