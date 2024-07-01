package golang

func countHomogenous(str string) int {
	var (
		curr     rune
		ans, cnt int
	)

	for _, r := range str {
		if r != curr {
			curr = r
			cnt = 0
		}

		cnt++

		ans = (ans + cnt) % MOD
	}

	return ans
}
