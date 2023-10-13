package robotWithString

import "strings"

type Stack struct {
	data []rune
}

func (s *Stack) Push(val rune) {
	s.data = append(s.data, val)
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Next() bool {
	return s.Len() > 0
}

func (s *Stack) Pop() rune {
	if !s.Next() {
		return -1
	}

	val := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]

	return val
}

func (s *Stack) Peek() rune {
	if !s.Next() {
		return -1
	}

	return s.data[s.Len()-1]
}

func min(a, b rune) rune {
	if a < b {
		return a
	}
	return b
}

func robotWithString(s string) string {
	mins := make([]rune, len(s))
	mina := 'z'
	for i := len(s) - 1; i >= 0; i-- {
		mina = min(mina, rune(s[i]))
		mins[i] = mina
	}

	i := 0
	stack := Stack{}
	sb := strings.Builder{}

	for i < len(s) {
		if !stack.Next() || stack.Peek() > mins[i] {
			stack.Push(rune(s[i]))
			i++
			continue
		}

		sb.WriteRune(stack.Pop())
	}

	for stack.Next() {
		sb.WriteRune(stack.Pop())
	}

	return sb.String()
}
