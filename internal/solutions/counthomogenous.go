package solutions

import "github.com/therenotomorrow/leetcode/internal/structs"

func countHomogenous(s string) int {
	var (
		curr     rune
		ans, cnt int
	)

	for _, r := range s {
		if r != curr {
			curr = r
			cnt = 0
		}

		cnt++

		ans = (ans + cnt) % structs.MOD
	}

	return ans
}
