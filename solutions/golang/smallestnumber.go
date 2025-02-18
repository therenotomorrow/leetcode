package golang

import (
	"strconv"
	"strings"
)

func smallestNumber(pattern string) string {
	ans := make([]string, 0)
	stack := NewStack[string]()

	for i, runa := range pattern + " " {
		stack.Push(strconv.Itoa(i + 1))

		if i == len(pattern) || runa == 'I' {
			for !stack.IsEmpty() {
				val, _ := stack.Pop()
				ans = append(ans, val)
			}
		}
	}

	return strings.Join(ans, "")
}
