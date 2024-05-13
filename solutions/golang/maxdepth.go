package golang

func maxDepth(s string) int {
	cnt := 0
	ans := 0

	for _, r := range s {
		if r == '(' {
			cnt++
		}

		if r == ')' {
			cnt--
		}

		ans = Max(ans, cnt)
	}

	return ans
}
