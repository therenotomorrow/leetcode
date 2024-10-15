package golang

type FirstUnique struct {
	queue Queue[int]
	count map[int]int
}

func FirstUniqueConstructor(nums []int) FirstUnique {
	uniq := FirstUnique{queue: NewQueue[int](), count: make(map[int]int)}

	for _, num := range nums {
		uniq.Add(num)
	}

	return uniq
}

func (fu *FirstUnique) ShowFirstUnique() int {
	for !fu.queue.IsEmpty() {
		peek, _ := fu.queue.Peek()

		if fu.count[peek] == 1 {
			return peek
		}

		_, _ = fu.queue.Dequeue()
	}

	return -1
}

func (fu *FirstUnique) Add(value int) {
	fu.count[value]++

	if fu.count[value] == 1 {
		fu.queue.Enqueue(value)
	}
}
