package golang

func backspace(s string, pos int) int {
	for skip := 0; pos >= 0; pos-- {
		switch {
		case s[pos] == '#':
			skip++
		case skip > 0:
			skip--
		default:
			return pos
		}
	}

	// string became empty
	return -1
}

func backspaceCompare(s string, t string) bool {
	i := len(s)
	j := len(t)

	for {
		i = backspace(s, i-1)
		j = backspace(t, j-1)

		if i < 0 || j < 0 {
			break
		}

		if s[i] != t[j] {
			return false
		}
	}

	return i == j
}
