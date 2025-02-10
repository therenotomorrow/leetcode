package golang

import (
	"slices"
	"unicode"
)

func clearDigits(s string) string {
	stack := NewStack[rune]()

	for _, runa := range s {
		if unicode.IsDigit(runa) {
			stack.Pop()
		} else {
			stack.Push(runa)
		}
	}

	ans := make([]rune, 0, stack.Size())

	for !stack.IsEmpty() {
		val, _ := stack.Pop()

		ans = append(ans, val)
	}

	slices.Reverse(ans)

	return string(ans)
}
