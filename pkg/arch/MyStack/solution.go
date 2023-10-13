package MyStack

type queue []int

func (q *queue) enqueue(val int) {
	*q = append(*q, val)
}

func (q *queue) peek() int {
	if q.empty() {
		return -1
	}

	return (*q)[0]
}

func (q *queue) dequeue() int {
	if q.empty() {
		return -1
	}

	val := q.peek()
	*q = (*q)[1:]

	return val
}

func (q *queue) len() int {
	return len(*q)
}

func (q *queue) empty() bool {
	return q.len() < 1
}

type MyStack struct {
	head queue
	tail queue
}

func Constructor() MyStack {
	return MyStack{head: make([]int, 0), tail: make([]int, 0)}
}

func (s *MyStack) Push(x int) {
	if s.head.empty() {
		s.head.enqueue(x)
		return
	}

	for !s.head.empty() {
		s.tail.enqueue(s.head.dequeue())
	}
	s.head.enqueue(x)
	for !s.tail.empty() {
		s.head.enqueue(s.tail.dequeue())
	}
}

func (s *MyStack) Pop() int {
	if s.head.empty() {
		return -1
	}

	return s.head.dequeue()
}

func (s *MyStack) Top() int {
	if s.head.empty() {
		return -1
	}

	return s.head.peek()
}

func (s *MyStack) Empty() bool {
	return s.head.empty() && s.tail.empty()
}
