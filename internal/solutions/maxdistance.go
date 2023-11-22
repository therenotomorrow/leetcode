package solutions

import (
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
	"math"
)

func maxDistance(arrays [][]int) int {
	lastMin := math.MaxInt
	currMin := math.MaxInt
	minIdx := 0

	lastMax := math.MinInt
	currMax := math.MinInt
	maxIdx := 0

	for i, arr := range arrays {
		switch {
		case arr[0] < currMin:
			lastMin = currMin
			currMin = arr[0]
			minIdx = i
		case arr[0] < lastMin:
			lastMin = arr[0]
		}

		switch n := len(arr) - 1; {
		case arr[n] > currMax:
			lastMax = currMax
			currMax = arr[n]
			maxIdx = i
		case arr[n] > lastMax:
			lastMax = arr[n]
		}
	}

	if minIdx == maxIdx {
		return mathfunc.Max(mathfunc.Abs(lastMax-currMin), mathfunc.Abs(currMax-lastMin))
	}

	return mathfunc.Max(mathfunc.Abs(currMax - currMin))
}
