package golang

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var (
		a, b, res int
		stack     = NewStack[int]()
	)

	for _, token := range tokens {
		if val, err := strconv.Atoi(token); err == nil {
			stack.Push(val)

			continue
		}

		a, _ = stack.Pop()
		b, _ = stack.Pop()

		switch token {
		case "+":
			res = b + a
		case "-":
			res = b - a
		case "*":
			res = b * a
		case "/":
			res = b / a
		}

		stack.Push(res)
	}

	val, _ := stack.Peek()

	return val
}
