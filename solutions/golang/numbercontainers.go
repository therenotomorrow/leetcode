package golang

import "container/heap"

type NumberContainers struct {
	indexes map[int]int
	// took from kthlargest.go: https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
	numbers map[int]*IntSlice
}

func NumberContainersConstructor() NumberContainers {
	return NumberContainers{
		indexes: make(map[int]int),
		numbers: make(map[int]*IntSlice),
	}
}

func (nc *NumberContainers) Change(index int, number int) {
	nc.indexes[index] = number

	if nc.numbers[number] == nil {
		nc.numbers[number] = new(IntSlice)
	}

	heap.Push(nc.numbers[number], index)
}

func (nc *NumberContainers) Find(number int) int {
	slice, ok := nc.numbers[number]
	if !ok {
		return -1
	}

	for slice.Len() > 0 {
		index := (*nc.numbers[number])[0]

		if nc.indexes[index] == number {
			return index
		}

		heap.Pop(nc.numbers[number])
	}

	return -1
}
