package golang

import "container/heap"

type IntSlice []int

func (ints *IntSlice) Len() int {
	return len(*ints)
}

func (ints *IntSlice) Less(i, j int) bool {
	return (*ints)[i] < (*ints)[j]
}

func (ints *IntSlice) Swap(i, j int) {
	(*ints)[i], (*ints)[j] = (*ints)[j], (*ints)[i]
}

func (ints *IntSlice) Push(x any) {
	val, ok := x.(int)
	if !ok {
		panic("wrong int")
	}

	*ints = append(*ints, val)
}

func (ints *IntSlice) Pop() any {
	old := *ints
	n := len(old)
	x := old[n-1]
	*ints = old[0 : n-1]

	return x
}

type KthLargest struct {
	k    int
	nums *IntSlice
}

func KthLargestConstructor(k int, nums []int) KthLargest {
	obj := KthLargest{k: k, nums: new(IntSlice)}

	for _, num := range nums {
		obj.Add(num)
	}

	return obj
}

func (kth *KthLargest) Add(val int) int {
	if kth.nums.Len() < kth.k || (*kth.nums)[0] < val {
		heap.Push(kth.nums, val)

		if kth.nums.Len() > kth.k {
			heap.Pop(kth.nums)
		}
	}

	return (*kth.nums)[0]
}
