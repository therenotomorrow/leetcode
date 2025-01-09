package golang

import (
	"sort"
)

func validSquare(point1 []int, point2 []int, point3 []int, point4 []int) bool {
	const (
		first = iota
		second
		third
		fourth
	)

	points := [][]int{point1, point2, point3, point4}
	distFunc := func(i, j int) int {
		xDiff := points[i][0] - points[j][0]
		yDiff := points[i][1] - points[j][1]

		return xDiff*xDiff + yDiff*yDiff
	}
	distances := []int{
		distFunc(first, second),
		distFunc(first, third),
		distFunc(first, fourth),
		distFunc(second, third),
		distFunc(second, fourth),
		distFunc(third, fourth),
	}

	sort.Ints(distances)

	return distances[0] != distances[5] && distances[0] == distances[3] && distances[4] == distances[5]
}
