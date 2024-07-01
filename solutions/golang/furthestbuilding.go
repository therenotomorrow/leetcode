package golang

import "container/heap"

func furthestBuilding(heights []int, bricks int, ladders int) int {
	peeks := new(IntHeap)

	heap.Init(peeks)

	for i := range len(heights) - 1 {
		diff := heights[i+1] - heights[i]

		if diff <= 0 {
			continue // don't need to climb
		}

		heap.Push(peeks, diff)

		if peeks.Len() <= ladders {
			continue // could use ladder
		}

		// could replace ladder with min diff with bricks
		diff, ok := heap.Pop(peeks).(int)

		if !ok {
			panic("wrong `diff`")
		}

		bricks -= diff

		if bricks < 0 {
			return i
		}
	}

	// we all in the end
	return len(heights) - 1
}

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x any) {
	val, ok := x.(int)

	if !ok {
		panic("wrong `x`")
	}

	*h = append(*h, val)
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
