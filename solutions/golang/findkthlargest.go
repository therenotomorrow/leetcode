package golang

import (
	"container/heap"
)

type Slice []int

func (ints *Slice) Len() int {
	return len(*ints)
}

func (ints *Slice) Less(i, j int) bool {
	return (*ints)[i] > (*ints)[j]
}

func (ints *Slice) Swap(i, j int) {
	(*ints)[i], (*ints)[j] = (*ints)[j], (*ints)[i]
}

func (ints *Slice) Push(x any) {
	val, ok := x.(int)
	if !ok {
		panic("wrong int")
	}

	*ints = append(*ints, val)
}

func (ints *Slice) Pop() any {
	old := *ints
	n := len(old)
	x := old[n-1]
	*ints = old[0 : n-1]

	return x
}

func findKthLargest(nums []int, k int) int {
	slice := new(Slice)

	for _, num := range nums {
		heap.Push(slice, num)
	}

	val := 0
	for ; k > 0; k-- {
		val, _ = heap.Pop(slice).(int)
	}

	return val
}
