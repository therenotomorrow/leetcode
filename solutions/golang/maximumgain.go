package golang

import "strings"

func maximumGain(text string, scoreX int, scoreY int) int {
	var (
		builder strings.Builder
		tmpl    = []rune{'a', 'b'}
		score   = 0
		order   = []int{scoreX, scoreY}
		stack   = NewStack[rune]()
	)

	if scoreY > scoreX {
		order[0], order[1] = scoreY, scoreX
		tmpl = []rune{'b', 'a'}
	}

	for _, point := range order {
		builder.Reset()

		for _, curr := range text {
			prev, _ := stack.Peek()

			if prev == tmpl[0] && curr == tmpl[1] {
				stack.Pop()

				score += point
			} else {
				stack.Push(curr)
			}
		}

		for !stack.IsEmpty() {
			val, _ := stack.Pop()
			builder.WriteRune(val)
		}

		text = builder.String()
	}

	return score
}
