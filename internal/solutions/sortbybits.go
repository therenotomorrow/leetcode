package solutions

import "sort"

func countBits(n int) int {
	// Brian Kernighan's algorithm
	cnt := 0

	for n > 0 {
		n &= n - 1
		cnt++
	}

	return cnt
}

func sortByBits(arr []int) []int {
	sort.SliceStable(arr, func(i, j int) bool {
		a := countBits(arr[i])
		b := countBits(arr[j])

		return a == b && arr[i] < arr[j] || a < b
	})

	return arr
}
