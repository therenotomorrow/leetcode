package golang

import "container/heap"

func makePrefSumNonNegative(nums []int) int {
	sum := 0
	cnt := 0
	// took from kthlargest.go: https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
	slice := new(IntSlice)

	for _, num := range nums {
		if num < 0 {
			heap.Push(slice, num)
		}

		sum += num

		if sum < 0 {
			val, ok := heap.Pop(slice).(int)
			if !ok {
				panic("bad int")
			}

			sum -= val
			cnt++
		}
	}

	return cnt
}
