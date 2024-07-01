package golang

func minTimeToVisitAllPoints(points [][]int) int {
	minTime := 0

	for i := range len(points) - 1 {
		fromX := points[i][0]
		fromY := points[i][1]

		toX := points[i+1][0]
		toY := points[i+1][1]

		minTime += Max(Abs(toX-fromX), Abs(toY-fromY))
	}

	return minTime
}
