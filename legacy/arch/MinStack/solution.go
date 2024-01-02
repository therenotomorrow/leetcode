package MinStack

type MyNode struct {
	val  int
	min  int
	next *MyNode
}

type MinStack struct {
	head *MyNode
}

func Constructor() MinStack {
	return MinStack{}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func (s *MinStack) Push(val int) {
	if s.head == nil {
		s.head = &MyNode{val: val, min: val}
	} else {
		s.head = &MyNode{val: val, min: min(val, s.head.min), next: s.head}
	}
}

func (s *MinStack) Pop() {
	s.head = s.head.next
}

func (s *MinStack) Top() int {
	return s.head.val
}

func (s *MinStack) GetMin() int {
	return s.head.min
}
