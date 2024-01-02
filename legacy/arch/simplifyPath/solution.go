package simplifyPath

import (
	"strings"
)

type Stack struct {
	data []string
	len  int
}

func (s *Stack) Push(val string) {
	s.data = append(s.data, val)
	s.len++
}

func (s *Stack) Pop() string {
	if s.len < 1 {
		return ""
	}

	val := s.data[s.len-1]
	s.data = s.data[:s.len-1]

	s.len--

	return val
}

func simplifyPath(path string) string {
	stack := Stack{}
	builder := strings.Builder{}

	for _, s := range strings.Split(path, "/") {
		if s == "." || s == "" {
			continue
		}

		if s == ".." {
			stack.Pop()
			continue
		}

		stack.Push(s)
	}

	for _, d := range stack.data {
		builder.WriteRune('/')
		builder.WriteString(d)
	}

	if len(stack.data) == 0 {
		return "/"
	}

	return builder.String()
}
