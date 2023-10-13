package backspaceCompare

func backspace(s string) string {
	stack := make([]rune, 0)

	for _, r := range s {
		switch l := len(stack); {
		case r != '#':
			stack = append(stack, r)
		case l > 0:
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}

func backspaceCompare(s string, t string) bool {
	return backspace(s) == backspace(t)
}
