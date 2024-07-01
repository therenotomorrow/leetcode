package totalCost

import "container/heap"

func totalCost(costs []int, k int, candidates int) int64 {
	minHeap := &MinHeap{}

	// Add first candidates to heap
	for i := 0; i < candidates; i++ {
		heap.Push(minHeap, &Item{cost: costs[i], pos: 0})
	}

	// Add last candidates to heap
	for i := maxi(candidates, len(costs)-candidates); i < len(costs); i++ {
		heap.Push(minHeap, &Item{cost: costs[i], pos: 1})
	}

	left, right := candidates, len(costs)-1-candidates
	res := 0

	for i := 0; i < k; i++ {
		candidate := heap.Pop(minHeap).(*Item)
		res += candidate.cost
		idx := candidate.pos

		if left <= right {
			if idx == 0 {
				heap.Push(minHeap, &Item{cost: costs[left], pos: 0})
				left++
			} else {
				heap.Push(minHeap, &Item{cost: costs[right], pos: 1})
				right--
			}
		}
	}

	return int64(res)
}

type Item struct {
	cost int
	pos  int
}

type MinHeap []*Item

func (h *MinHeap) Len() int { return len(*h) }

func (h *MinHeap) Less(i, j int) bool {
	if (*h)[i].cost == (*h)[j].cost {
		return (*h)[i].pos < (*h)[j].pos
	}
	return (*h)[i].cost < (*h)[j].cost
}

func (h *MinHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*Item))
}

func (h *MinHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}
