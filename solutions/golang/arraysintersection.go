package golang

import (
	"slices"
)

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	const numOfArrays = 3

	cnt := make(map[int]int)

	for _, num := range slices.Concat(arr1, arr2, arr3) {
		cnt[num]++
	}

	arr := make([]int, 0)

	for _, num := range arr1 {
		if cnt[num] == numOfArrays {
			arr = append(arr, num)
		}
	}

	return arr
}
