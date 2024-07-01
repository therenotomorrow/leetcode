package golang

import (
	"sort"
)

func assignBikes1(workers [][]int, bikes [][]int) []int {
	distances := make([][]int, 0)

	for i, worker := range workers {
		for j, bike := range bikes {
			distances = append(distances, []int{Manhattan(worker[0], worker[1], bike[0], bike[1]), i, j})
		}
	}

	sort.SliceStable(distances, func(i, j int) bool {
		one := distances[i]
		two := distances[j]

		// sort by distance
		if one[0] != two[0] {
			return one[0] < two[0]
		}

		// sort by worker number
		if one[1] != two[1] {
			return one[1] < two[1]
		}

		return one[2] < two[2]
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
