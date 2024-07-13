package survivedRobotsHealths

import (
	"slices"
)

type Stackable interface {
	~int | ~rune
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

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	n := len(positions)
	indices := make([]int, n)
	result := make([]int, 0)
	stack := NewStack[int]()

	for index := range n {
		indices[index] = index
	}

	slices.SortStableFunc(indices, func(a, b int) int {
		return positions[a] - positions[b]
	})

	for _, currentIndex := range indices {
		// Add right-moving robots to the stack
		if directions[currentIndex] == 'R' {
			stack.Push(currentIndex)
		} else {
			for !stack.IsEmpty() && healths[currentIndex] > 0 {
				// Pop the top robot from the stack for collision check
				topIndex, _ := stack.Pop()

				// Top robot survives, current robot is destroyed
				if healths[topIndex] > healths[currentIndex] {
					healths[topIndex] -= 1
					healths[currentIndex] = 0
					stack.Push(topIndex)
				} else if healths[topIndex] < healths[currentIndex] {
					// Current robot survives, top robot is destroyed
					healths[currentIndex] -= 1
					healths[topIndex] = 0
				} else {
					// Both robots are destroyed
					healths[currentIndex] = 0
					healths[topIndex] = 0
				}
			}
		}
	}

	// Collect surviving robots
	for index := range n {
		if healths[index] > 0 {
			result = append(result, healths[index])
		}
	}
	return result
}
