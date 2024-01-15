package solutions

import (
	"sort"
	"time"

	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func findMinDifference(timePoints []string) int {
	const (
		format   = "15:04"
		totalMin = 24.0 * 60
	)

	sort.Strings(timePoints)

	diff := totalMin

	for i := 0; i < len(timePoints)-1; i++ {
		prev, _ := time.Parse(format, timePoints[i])
		curr, _ := time.Parse(format, timePoints[i+1])

		diff = mathfunc.Min(diff, prev.Sub(curr).Abs().Minutes())
	}

	first, _ := time.Parse(format, timePoints[0])
	last, _ := time.Parse(format, timePoints[len(timePoints)-1])

	return int(mathfunc.Min(diff, first.Add(time.Minute*totalMin).Sub(last).Abs().Minutes()))
}
