package golang

func countSubstrings(str string) int {
	var (
		visited = make([][]bool, len(str))
		dynamic func(i int, j int) int
	)

	for i := range visited {
		visited[i] = make([]bool, len(str))
	}

	dynamic = func(i int, j int) int {
		if i > j || j < 0 || i >= len(str) || visited[i][j] {
			return 0
		}

		palindrome := 1

		for left, right := i, j; left <= right; left, right = left+1, right-1 {
			if str[left] != str[right] {
				palindrome = 0

				break
			}
		}

		palindrome += dynamic(i, j-1) + dynamic(i+1, j)

		visited[i][j] = true

		return palindrome
	}

	return dynamic(0, len(str)-1)
}
