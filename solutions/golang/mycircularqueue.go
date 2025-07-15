package golang

type MyCircularQueue struct {
	head     int
	maxSize  int
	currSize int
	data     []int
}

func MyCircularQueueConstructor(k int) MyCircularQueue {
	return MyCircularQueue{head: 0, maxSize: k, currSize: 0, data: make([]int, k)}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}

	q.data[(q.head+q.currSize)%q.maxSize] = value

	q.currSize++

	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	q.head = (q.head + 1) % q.maxSize

	q.currSize--

	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[q.head]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[(q.head+q.currSize-1)%q.maxSize]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.currSize == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return q.currSize == q.maxSize
}
