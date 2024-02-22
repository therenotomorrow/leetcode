package golang

import "container/heap"

type SticksPQ []int

func (pq *SticksPQ) Len() int {
	return len(*pq)
}

func (pq *SticksPQ) Less(i, j int) bool {
	return (*pq)[i] < (*pq)[j]
}

func (pq *SticksPQ) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *SticksPQ) Push(x interface{}) {
	val, ok := x.(int)

	if !ok {
		panic("wrong int")
	}

	*pq = append(*pq, val)
}

func (pq *SticksPQ) Pop() interface{} {
	last := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]

	return last
}

func connectSticks(sticks []int) int {
	pq := SticksPQ(sticks)
	sum := 0

	heap.Init(&pq)

	for pq.Len() > 1 {
		a, ok := heap.Pop(&pq).(int)
		if !ok {
			panic("wrong `a`")
		}

		b, ok := heap.Pop(&pq).(int)
		if !ok {
			panic("wrong `b`")
		}

		sum += a + b

		heap.Push(&pq, a+b)
	}

	return sum
}
