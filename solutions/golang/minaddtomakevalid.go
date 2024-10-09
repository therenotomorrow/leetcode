package golang

func minAddToMakeValid(s string) int {
	var stack, moves int

	for _, r := range s {
		switch {
		case r == '(':
			stack++
		case stack > 0:
			stack--
		default:
			moves++
		}
	}

	return stack + moves
}
