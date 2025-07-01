package golang

import "strings"

func robotWithString(text string) string {
	mins := make([]rune, len(text))
	curr := 'z'

	for i := len(text) - 1; i >= 0; i-- {
		curr = Min(curr, rune(text[i]))
		mins[i] = curr
	}

	stack := NewStack[rune]()
	printed := strings.Builder{}

	for i := 0; i < len(text); {
		if peek, _ := stack.Peek(); stack.IsEmpty() || peek > mins[i] {
			stack.Push(rune(text[i]))

			i++

			continue
		}

		val, _ := stack.Pop()

		printed.WriteRune(val)
	}

	for !stack.IsEmpty() {
		val, _ := stack.Pop()

		printed.WriteRune(val)
	}

	return printed.String()
}
