package MyQueue

type Node struct {
	val  int
	next *Node
}

type Stack struct {
	head *Node
	len  int
}

func (s *Stack) Push(val int) {
	s.head = &Node{val: val, next: s.head}
	s.len++
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}

	return s.head.val
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	val := s.head.val
	s.head = s.head.next
	s.len--

	return val
}

func (s *Stack) Size() int {
	return s.len
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

type MyQueue struct {
	head *Stack
	tail *Stack
}

func Constructor() MyQueue {
	return MyQueue{head: &Stack{}, tail: &Stack{}}
}

func (q *MyQueue) swap() {
	for !q.tail.IsEmpty() {
		q.head.Push(q.tail.Pop())
	}
}

func (q *MyQueue) Push(x int) {
	q.tail.head = &Node{val: x, next: q.tail.head}
	q.tail.len++
}

func (q *MyQueue) Pop() int {
	if !q.head.IsEmpty() {
		return q.head.Pop()
	}

	q.swap()

	return q.head.Pop()
}

func (q *MyQueue) Peek() int {
	if !q.head.IsEmpty() {
		return q.head.Peek()
	}

	q.swap()

	return q.head.Peek()
}

func (q *MyQueue) Empty() bool {
	return q.head.len+q.tail.len == 0
}
