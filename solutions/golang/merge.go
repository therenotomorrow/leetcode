package golang

import "sort"

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prev := intervals[0]
	ans := [][]int{prev}

	for _, interval := range intervals {
		if prev[1] < interval[0] {
			ans = append(ans, interval)
			prev = interval
		} else {
			prev[1] = Max(prev[1], interval[1])
		}
	}

	return ans
}
