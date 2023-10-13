package makeGood

func abs(a rune) rune {
	if a < 0 {
		return -a
	}

	return a
}

func makeGood(s string) string {
	stack := make([]rune, 0)

	for _, r := range s {
		if l := len(stack); l != 0 && abs(r-stack[l-1]) == 32 {
			stack = stack[:l-1]
		} else {
			stack = append(stack, r)
		}
	}

	return string(stack)
}
