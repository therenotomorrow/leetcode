package golang

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	xArr := make([]int, 0)
	for _, point := range points {
		xArr = append(xArr, point[0])
	}

	sort.Ints(xArr)

	maxWidth := xArr[1] - xArr[0]
	for i := 2; i < len(xArr); i++ {
		if xArr[i]-xArr[i-1] > maxWidth {
			maxWidth = xArr[i] - xArr[i-1]
		}
	}

	return maxWidth
}
