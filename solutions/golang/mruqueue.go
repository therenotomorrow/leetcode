package golang

import "container/list"

type MRUQueue struct {
	data *list.List
}

func MRUQueueConstructor(n int) MRUQueue {
	data := list.New()

	for i := 1; i <= n; i++ {
		data.PushBack(i)
	}

	return MRUQueue{data: data}
}

func (q *MRUQueue) Fetch(k int) int {
	curr := q.data.Front()
	for i := 1; i < k; i++ {
		curr = curr.Next()
	}

	val, ok := curr.Value.(int)
	if !ok {
		panic("wrong int")
	}

	q.data.Remove(curr)
	q.data.PushBack(val)

	return val
}
