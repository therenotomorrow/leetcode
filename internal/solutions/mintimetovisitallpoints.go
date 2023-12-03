package solutions

import "github.com/therenotomorrow/leetcode/pkg/mathfunc"

func minTimeToVisitAllPoints(points [][]int) int {
	minTime := 0

	for i := 0; i < len(points)-1; i++ {
		fromX := points[i][0]
		fromY := points[i][1]

		toX := points[i+1][0]
		toY := points[i+1][1]

		minTime += mathfunc.Max(mathfunc.Abs(toX-fromX), mathfunc.Abs(toY-fromY))
	}

	return minTime
}
