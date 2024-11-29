package golang

import (
	"math"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	overlap := 0
	start := math.MinInt

	for _, interval := range intervals {
		x := interval[0]
		y := interval[1]

		if x >= start {
			start = y
		} else {
			overlap++
		}
	}

	return overlap
}
