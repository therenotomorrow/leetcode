package datastruct

import "github.com/therenotomorrow/leetcode/internal/structs"

type Queueable interface {
	~int | ~*structs.TreeNode
}

type Queue[T Queueable] interface {
	Enqueue(val T)
	Dequeue() (val T, ok bool)
	Peek() (val T, ok bool)
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
	return &queue[T]{}
}

func (q *queue[T]) Enqueue(val T) {
	old := q.tail
	q.tail = &qNode[T]{val: val}

	if q.IsEmpty() {
		q.head = q.tail
	} else {
		old.next = q.tail
	}

	q.size++
}

func (q *queue[T]) Dequeue() (val T, ok bool) {
	if q.IsEmpty() {
		return
	}

	val = q.head.val
	q.head = q.head.next
	q.size--

	return val, true
}

func (q *queue[T]) Peek() (val T, ok bool) {
	if q.IsEmpty() {
		return
	}

	return q.head.val, true
}

func (q *queue[T]) Size() int {
	return q.size
}

func (q *queue[T]) IsEmpty() bool {
	return q.Size() == 0
}
