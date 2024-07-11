package golang

import "strings"

func reverseParentheses(s string) string {
	stack := NewStack[int]()
	runes := []rune(s)

	for j, runa := range runes {
		switch runa {
		case '(':
			stack.Push(j + 1)
		case ')':
			i, _ := stack.Pop()

			for i < j {
				runes[i], runes[j] = runes[j], runes[i]
				i++
				j--
			}
		}
	}

	var builder strings.Builder

	for _, runa := range runes {
		switch runa {
		case '(', ')':
			continue
		default:
			builder.WriteRune(runa)
		}
	}

	return builder.String()
}
