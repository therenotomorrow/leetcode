package golang

import "sort"

func canBeEqual(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)

	for i := range target {
		if target[i] != arr[i] {
			return false
		}
	}

	return true
}
