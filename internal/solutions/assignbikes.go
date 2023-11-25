package solutions

import (
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
	"sort"
)

func manhattanDistance(x1 int, y1 int, x2 int, y2 int) int {
	return mathfunc.Abs(x1-x2) + mathfunc.Abs(y1-y2)
}

func assignBikes(workers [][]int, bikes [][]int) []int {
	distances := make([][]int, 0)

	for i, worker := range workers {
		for j, bike := range bikes {
			distances = append(distances, []int{manhattanDistance(worker[0], worker[1], bike[0], bike[1]), i, j})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		a := distances[i]
		b := distances[j]

		// sort by distance
		if a[0] != b[0] {
			return a[0] < b[0]
		}

		// sort by worker number
		if a[1] != b[1] {
			return a[1] < b[1]
		}

		return a[2] < b[2]
	})

	assigned := make([]int, len(workers))
	for i := range assigned {
		assigned[i] = -1
	}
	taken := make([]bool, len(bikes))

	for _, distance := range distances {
		worker := distance[1]
		bike := distance[2]

		if assigned[worker] != -1 || taken[bike] {
			continue
		}

		taken[bike] = true
		assigned[worker] = bike
	}

	return assigned
}
