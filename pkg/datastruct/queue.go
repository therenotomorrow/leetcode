package datastruct

type Queue interface {
	Enqueue(val int)
	Dequeue() (val int, ok bool)
	Peek() (val int, ok bool)
	Size() int
	IsEmpty() bool
}

type qNode struct {
	val  int
	next *qNode
}

type queue struct {
	head *qNode
	tail *qNode
	size int
}

func NewQueue() Queue {
	return &queue{}
}

func (q *queue) Enqueue(val int) {
	old := q.tail
	q.tail = &qNode{val: val}

	if q.Size() == 0 {
		q.head = q.tail
	} else {
		old.next = q.tail
	}

	q.size++
}

func (q *queue) Dequeue() (val int, ok bool) {
	if q.Size() == 0 {
		return
	}

	val = q.head.val
	q.head = q.head.next
	q.size--

	return val, true
}

func (q *queue) Peek() (val int, ok bool) {
	if q.Size() == 0 {
		return
	}

	return q.head.val, true
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) IsEmpty() bool {
	return q.Size() == 0
}
