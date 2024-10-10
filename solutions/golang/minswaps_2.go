package golang

func minSwaps2(s string) int {
	const half = 2

	stack := NewStack[rune]()

	for _, runa := range s {
		switch runa {
		case '[':
			stack.Push(runa)
		default:
			_, _ = stack.Pop()
		}
	}

	return (stack.Size() + 1) / half
}