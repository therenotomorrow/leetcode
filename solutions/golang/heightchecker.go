package golang

import "sort"

func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))

	copy(sorted, heights)
	sort.Ints(sorted)

	ans := 0

	for i, h := range heights {
		if h != sorted[i] {
			ans++
		}
	}

	return ans
}
