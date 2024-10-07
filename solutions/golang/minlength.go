package golang

func minLength(s string) int {
	stack := NewStack[rune]()

	for _, curr := range s {
		last, _ := stack.Peek()

		if last == 'A' && curr == 'B' || last == 'C' && curr == 'D' {
			_, _ = stack.Pop()
		} else {
			stack.Push(curr)
		}
	}

	return stack.Size()
}
