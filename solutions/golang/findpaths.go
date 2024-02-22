package golang

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	var (
		c       = NewCache()
		dynamic func(currMove int, i int, j int, paths int) int
	)

	dynamic = func(currMove int, i int, j int, paths int) int {
		if i < 0 || j < 0 || i == m || j == n {
			return 1
		}

		if currMove == 0 {
			return 0
		}

		if val, ok := c.Load(i, j, currMove); ok {
			return val
		}

		paths += dynamic(currMove-1, i-1, j, paths) +
			dynamic(currMove-1, i+1, j, paths) +
			dynamic(currMove-1, i, j-1, paths) +
			dynamic(currMove-1, i, j+1, paths)
		paths %= MOD

		c.Save(paths, i, j, currMove)

		return paths
	}

	return dynamic(maxMove, startRow, startColumn, 0)
}
