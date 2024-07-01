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

func (pq *SticksPQ) Push(x any) {
	val, ok := x.(int)

	if !ok {
		panic("wrong int")
	}

	*pq = append(*pq, val)
}

func (pq *SticksPQ) Pop() any {
	last := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]

	return last
}

func connectSticks(sticks []int) int {
	pQue := SticksPQ(sticks)
	sum := 0

	heap.Init(&pQue)

	for pQue.Len() > 1 {
		one, exist := heap.Pop(&pQue).(int)
		if !exist {
			panic("wrong `one`")
		}

		two, exist := heap.Pop(&pQue).(int)
		if !exist {
			panic("wrong `two`")
		}

		sum += one + two

		heap.Push(&pQue, one+two)
	}

	return sum
}
