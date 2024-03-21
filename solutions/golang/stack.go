package golang

type Stackable interface {
	~int | ~rune | ~*TreeNode | ~*ListNode
}

type Stack[T Stackable] interface {
	Push(val T)
	Pop() (T, bool)
	Peek() (T, bool)
	Size() int
	IsEmpty() bool
}

type sNode[T Stackable] struct {
	val  T
	next *sNode[T]
}

type stack[T Stackable] struct {
	head *sNode[T]
	size int
}

func NewStack[T Stackable]() Stack[T] {
	return &stack[T]{head: nil, size: 0}
}

func (s *stack[T]) Push(val T) {
	s.head = &sNode[T]{val: val, next: s.head}
	s.size++
}

func (s *stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T

		return zero, false
	}

	val := s.head.val
	s.head = s.head.next
	s.size--

	return val, true
}

func (s *stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T

		return zero, false
	}

	return s.head.val, true
}

func (s *stack[T]) Size() int {
	return s.size
}

func (s *stack[T]) IsEmpty() bool {
	return s.Size() == 0
}
