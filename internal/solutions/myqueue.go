package solutions

import "github.com/therenotomorrow/leetcode/pkg/datastruct"

type MyQueue struct {
	head datastruct.Stack[int]
	tail datastruct.Stack[int]
}

func MyQueueConstructor() MyQueue {
	return MyQueue{
		head: datastruct.NewStack[int](),
		tail: datastruct.NewStack[int](),
	}
}

func (q *MyQueue) swap() {
	for !q.tail.IsEmpty() {
		val, _ := q.tail.Pop()
		q.head.Push(val)
	}
}

func (q *MyQueue) Push(x int) {
	q.tail.Push(x)
}

func (q *MyQueue) Pop() int {
	if !q.head.IsEmpty() {
		val, _ := q.head.Pop()

		return val
	}

	q.swap()

	val, _ := q.head.Pop()

	return val
}

func (q *MyQueue) Peek() int {
	if !q.head.IsEmpty() {
		val, _ := q.head.Peek()

		return val
	}

	q.swap()

	val, _ := q.head.Peek()

	return val
}

func (q *MyQueue) Empty() bool {
	return q.head.IsEmpty() && q.tail.IsEmpty()
}
