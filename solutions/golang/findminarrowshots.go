package golang

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	curr := points[0][1]
	arrows := 1 // min possible arrows

	for _, point := range points {
		start := point[0]
		end := point[1]

		if curr < start {
			curr = end
			arrows++
		}
	}

	return arrows
}
