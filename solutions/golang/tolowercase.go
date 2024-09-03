package golang

func toLowerCase(s string) string {
	ans := make([]rune, len(s))

	for i, r := range s {
		if 'A' <= r && r <= 'Z' {
			r += 32
		}

		ans[i] = r
	}

	return string(ans)
}
