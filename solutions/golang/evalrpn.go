package golang

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var (
		one, two, res int
		stack         = NewStack[int]()
	)

	for _, token := range tokens {
		if val, err := strconv.Atoi(token); err == nil {
			stack.Push(val)

			continue
		}

		one, _ = stack.Pop()
		two, _ = stack.Pop()

		switch token {
		case "+":
			res = two + one
		case "-":
			res = two - one
		case "*":
			res = two * one
		case "/":
			res = two / one
		}

		stack.Push(res)
	}

	val, _ := stack.Peek()

	return val
}
