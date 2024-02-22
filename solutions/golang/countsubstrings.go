package golang

func countSubstrings(s string) int {
	var (
		visited = make([][]bool, len(s))
		dynamic func(i int, j int) int
	)

	for i := range visited {
		visited[i] = make([]bool, len(s))
	}

	dynamic = func(i int, j int) int {
		if i > j || j < 0 || i >= len(s) || visited[i][j] {
			return 0
		}

		palindrome := 1

		for left, right := i, j; left <= right; left, right = left+1, right-1 {
			if s[left] != s[right] {
				palindrome = 0

				break
			}
		}

		palindrome += dynamic(i, j-1) + dynamic(i+1, j)

		visited[i][j] = true

		return palindrome
	}

	return dynamic(0, len(s)-1)
}
