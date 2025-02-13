package golang

import "container/heap"

func minOperations6(nums []int, k int) int {
	cnt := 0
	// took from kthlargest.go: https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
	slice := new(IntSlice)

	for _, num := range nums {
		heap.Push(slice, num)
	}

	for slice.Len() > 1 {
		one, isOne := heap.Pop(slice).(int)
		if !isOne {
			panic("bad int")
		}

		if one >= k {
			break
		}

		two, isTwo := heap.Pop(slice).(int)
		if !isTwo {
			panic("bad int")
		}

		cnt++

		heap.Push(slice, 2*Min(one, two)+Max(one, two))
	}

	return cnt
}
