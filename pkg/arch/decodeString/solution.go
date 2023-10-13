package decodeString

type stack []rune

func (s *stack) len() int {
	return len(*s)
}

func (s *stack) next() bool {
	return s.len() > 0
}

func (s *stack) push(val rune) {
	*s = append(*s, val)
}

func (s *stack) peek() rune {
	if !s.next() {
		return -1
	}

	return (*s)[s.len()-1]
}

func (s *stack) pop() rune {
	if !s.next() {
		return -1
	}

	val := s.peek()
	*s = (*s)[:s.len()-1]

	return val
}

func isDigit(val rune) bool {
	return '0' <= val && val <= '9'
}

func decodeString(s string) string {
	st := stack{}
	rs := make([]rune, 0)

	var (
		k int32 = 0
		b int32 = 1
	)

	for _, r := range s {
		switch r {
		case ']':
			for st.peek() != '[' {
				rs = append(rs, st.pop())
			}

			st.pop() // for '['

			for st.next() && isDigit(st.peek()) {
				k += (st.pop() - '0') * b
				b *= 10
			}

			for ; k != 0; k-- {
				for i := len(rs) - 1; i >= 0; i-- {
					st.push(rs[i])
				}
			}

			b = 1
			rs = nil
		default:
			st.push(r)
		}
	}

	return string(st)
}
