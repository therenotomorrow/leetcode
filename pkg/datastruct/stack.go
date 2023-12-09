package datastruct

import "github.com/therenotomorrow/leetcode/internal/structs"

type Stackable interface {
	~int | ~*structs.TreeNode
}

type Stack[T Stackable] interface {
	Push(val T)
	Pop() (val T, ok bool)
	Peek() (val T, ok bool)
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
	return &stack[T]{}
}

func (s *stack[T]) Push(val T) {
	s.head = &sNode[T]{val: val, next: s.head}
	s.size++
}

func (s *stack[T]) Pop() (val T, ok bool) {
	if s.IsEmpty() {
		return
	}

	val = s.head.val
	s.head = s.head.next
	s.size--

	return val, true
}

func (s *stack[T]) Peek() (val T, ok bool) {
	if s.IsEmpty() {
		return
	}

	return s.head.val, true
}

func (s *stack[T]) Size() int {
	return s.size
}

func (s *stack[T]) IsEmpty() bool {
	return s.Size() == 0
}
