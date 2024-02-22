package golang

type Queueable interface {
	~int | ~*TreeNode
}

type Queue[T Queueable] interface {
	Enqueue(val T)
	Dequeue() (T, bool)
	Peek() (T, bool)
	Size() int
	IsEmpty() bool
}

type qNode[T Queueable] struct {
	val  T
	next *qNode[T]
}

type queue[T Queueable] struct {
	head *qNode[T]
	tail *qNode[T]
	size int
}

func NewQueue[T Queueable]() Queue[T] {
	return &queue[T]{head: nil, tail: nil, size: 0}
}

func (q *queue[T]) Enqueue(val T) {
	old := q.tail
	q.tail = &qNode[T]{val: val, next: nil}

	if q.IsEmpty() {
		q.head = q.tail
	} else {
		old.next = q.tail
	}

	q.size++
}

func (q *queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T

		return zero, false
	}

	val := q.head.val
	q.head = q.head.next
	q.size--

	return val, true
}

func (q *queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T

		return zero, false
	}

	return q.head.val, true
}

func (q *queue[T]) Size() int {
	return q.size
}

func (q *queue[T]) IsEmpty() bool {
	return q.Size() == 0
}
