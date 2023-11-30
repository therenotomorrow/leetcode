package solutions

import "github.com/therenotomorrow/leetcode/internal/structs"

func countHomogenous(s string) int {
	curr := rune(0)
	ans := 0
	cnt := 0

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
