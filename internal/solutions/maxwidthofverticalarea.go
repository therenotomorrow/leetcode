package solutions

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	x := make([]int, 0)
	for _, point := range points {
		x = append(x, point[0])
	}

	sort.Ints(x)

	maxWidth := x[1] - x[0]
	for i := 2; i < len(x); i++ {
		if x[i]-x[i-1] > maxWidth {
			maxWidth = x[i] - x[i-1]
		}
	}

	return maxWidth
}
