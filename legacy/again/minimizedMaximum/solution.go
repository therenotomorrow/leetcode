package minimizedMaximum

import (
	"slices"
)

func canDistribute(x int, quantities []int, n int) bool {
	i := 0
	remaining := quantities[i]

	for range n {
		if remaining > x {
			remaining -= x

			continue
		}

		i++

		if i == len(quantities) {
			return true
		}

		remaining = quantities[i]
	}

	return false
}

func minimizedMaximum(n int, quantities []int) int {
	left := 0
	right := slices.Max(quantities)

	for left < right {
		middle := left + (right-left)/2

		if canDistribute(middle, quantities, n) {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return left
}
