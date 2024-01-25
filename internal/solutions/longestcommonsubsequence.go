package solutions

import "github.com/therenotomorrow/leetcode/pkg/mathfunc"

func longestCommonSubsequence(text1 string, text2 string) int {
	var (
		memo    = make([][]int, len(text1)+1)
		dynamic func(i int, j int) int
	)

	for i := range memo {
		memo[i] = make([]int, len(text2)+1)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	dynamic = func(i int, j int) int {
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if i == len(text1) || j == len(text2) {
			return 0
		}

		if text1[i] == text2[j] {
			memo[i][j] = 1 + dynamic(i+1, j+1)
		} else {
			memo[i][j] = mathfunc.Max(dynamic(i+1, j), dynamic(i, j+1))
		}

		return memo[i][j]
	}

	return dynamic(0, 0)
}
