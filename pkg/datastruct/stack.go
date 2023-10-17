package datastruct

type Stack interface {
	Push(val int)
	Pop() (val int, ok bool)
	Peek() (val int, ok bool)
	Size() int
	IsEmpty() bool
}

type sNode struct {
	val  int
	next *sNode
}

type stack struct {
	head *sNode
	size int
}

func NewStack() Stack {
	return &stack{}
}

func (s *stack) Push(val int) {
	s.head = &sNode{val: val, next: s.head}
	s.size++
}

func (s *stack) Pop() (val int, ok bool) {
	if s.Size() == 0 {
		return
	}

	val = s.head.val
	s.head = s.head.next
	s.size--

	return val, true
}

func (s *stack) Peek() (val int, ok bool) {
	if s.Size() == 0 {
		return
	}

	return s.head.val, true
}

func (s *stack) Size() int {
	return s.size
}

func (s *stack) IsEmpty() bool {
	return s.Size() == 0
}
