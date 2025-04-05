package golang

func minimumChairs(s string) int {
	ans := 0
	curr := 0

	for _, r := range s {
		if r == 'E' {
			curr++
		} else {
			curr--
		}

		ans = Max(ans, curr)
	}

	return ans
}
