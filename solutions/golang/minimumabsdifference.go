package golang

import (
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	minDiff := arr[1] - arr[0]
	ans := make([][]int, 0)

	for i := 1; i < len(arr); i++ {
		pair := []int{arr[i-1], arr[i]}

		switch currDiff := arr[i] - arr[i-1]; {
		case currDiff < minDiff:
			minDiff = currDiff

			ans = [][]int{pair}

		case currDiff == minDiff:
			ans = append(ans, pair)
		}
	}

	return ans
}
