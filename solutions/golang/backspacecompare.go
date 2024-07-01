package golang

func backspace(str string, pos int) int {
	for skip := 0; pos >= 0; pos-- {
		switch {
		case str[pos] == '#':
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

func backspaceCompare(str string, tpl string) bool {
	i := len(str)
	j := len(tpl)

	for {
		i = backspace(str, i-1)
		j = backspace(tpl, j-1)

		if i < 0 || j < 0 {
			break
		}

		if str[i] != tpl[j] {
			return false
		}
	}

	return i == j
}
