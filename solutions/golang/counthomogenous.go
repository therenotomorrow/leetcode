package golang

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

		ans = (ans + cnt) % MOD
	}

	return ans
}
