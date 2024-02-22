package golang

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

		ans = (ans + cnt) % MOD
	}

	return ans
}
