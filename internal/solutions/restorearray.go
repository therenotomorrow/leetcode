package solutions

import "github.com/therenotomorrow/leetcode/pkg/datastruct"

func restoreArray(adjacentPairs [][]int) []int {
	graph := make(map[int][]int)

	for _, pair := range adjacentPairs {
		graph[pair[0]] = append(graph[pair[0]], pair[1])
		graph[pair[1]] = append(graph[pair[1]], pair[0])
	}

	edge := 0
	for num, neighbors := range graph {
		if len(neighbors) == 1 {
			edge = num
			break
		}
	}

	ans := make([]int, len(graph))
	used := datastruct.NewSet[int]()

	ans[0] = edge
	for i := 1; i < len(ans); i++ {
		for _, neighbor := range graph[edge] {
			if used.Contains(neighbor) {
				continue
			}

			used.Add(edge)
			ans[i] = neighbor
			edge = neighbor
		}
	}

	return ans
}
