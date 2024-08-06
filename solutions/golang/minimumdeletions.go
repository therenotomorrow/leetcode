package golang

func minimumDeletions(text string) int {
	ans := 0
	stack := NewStack[rune]()

	for _, runa := range text {
		peek, ok := stack.Peek()
		if ok && peek == 'b' && runa == 'a' {
			ans++

			stack.Pop()
		} else {
			stack.Push(runa)
		}
	}

	return ans
}
