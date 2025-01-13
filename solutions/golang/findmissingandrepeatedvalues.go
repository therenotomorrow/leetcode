package golang

import "slices"

func findMissingAndRepeatedValues(grid [][]int) []int {
	seen := NewSet[int]()
	size := 0
	ans := []int{0, 0}

	for _, num := range slices.Concat(grid...) {
		if seen.Contains(num) {
			ans[0] = num
		} else {
			seen.Add(num)
		}

		size++
	}

	for num := 1; num <= size; num++ {
		if !seen.Contains(num) {
			ans[1] = num

			break
		}
	}

	return ans
}
