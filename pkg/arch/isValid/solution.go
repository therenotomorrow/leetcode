package isValid

type Node struct {
	next *Node
	val  rune
}

type Stack struct {
	head *Node
}

func (s *Stack) Pop() (rune, bool) {
	if s.head == nil {
		return 0, false
	}

	n := s.head
	s.head = n.next

	return n.val, true
}

func (s *Stack) Push(val rune) {
	n := &Node{val: val}

	s.head, n.next = n, s.head
}

func isValid(s string) bool {
	stack := Stack{}

	for _, r := range s {
		switch r {
		case '[', '(', '{':
			stack.Push(r)
		case ']', ')', '}':
			val, exist := stack.Pop()
			if !exist {
				return false
			}

			switch {
			case r == ']' && val == '[':
			case r == ')' && val == '(':
			case r == '}' && val == '{':
			default:
				return false
			}
		}
	}

	if stack.head != nil {
		return false
	}

	return true
}
