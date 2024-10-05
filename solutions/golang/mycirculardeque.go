package golang

type MyCircularDeque struct {
	size int
	data []int
}

func MyCircularDequeConstructor(k int) MyCircularDeque {
	return MyCircularDeque{size: k, data: make([]int, 0)}
}

func (deq *MyCircularDeque) InsertFront(value int) bool {
	if deq.IsFull() {
		return false
	}

	// need to use better indexes overlapping - I know about this :)
	deq.data = append([]int{value}, deq.data...)

	return true
}

func (deq *MyCircularDeque) InsertLast(value int) bool {
	if deq.IsFull() {
		return false
	}

	deq.data = append(deq.data, value)

	return true
}

func (deq *MyCircularDeque) DeleteFront() bool {
	if deq.IsEmpty() {
		return false
	}

	front := 1
	deq.data = deq.data[front:]

	return true
}

func (deq *MyCircularDeque) DeleteLast() bool {
	if deq.IsEmpty() {
		return false
	}

	last := len(deq.data) - 1
	deq.data = deq.data[:last]

	return true
}

func (deq *MyCircularDeque) GetFront() int {
	if deq.IsEmpty() {
		return -1
	}

	front := 0

	return deq.data[front]
}

func (deq *MyCircularDeque) GetRear() int {
	if deq.IsEmpty() {
		return -1
	}

	last := len(deq.data) - 1

	return deq.data[last]
}

func (deq *MyCircularDeque) IsEmpty() bool {
	return len(deq.data) == 0
}

func (deq *MyCircularDeque) IsFull() bool {
	return len(deq.data) == deq.size
}
