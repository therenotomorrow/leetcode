package treeDiameter

import (
	"cmp"
)

func treeDiameter(edges [][]int) int {
	var (
		graph   = make(map[int][]int)
		visited = make([]bool, len(edges)+1)
		dfs     func(node int) int
	)

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	diameter := 0

	dfs = func(node int) int {
		top1 := 0
		top2 := 0
		distance := 0
		visited[node] = true

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				distance = 1 + dfs(neighbor)
			}

			switch {
			case distance > top1:
				top1, top2 = distance, top1
			case distance > top2:
				top2 = distance
			}
		}

		diameter = Max(diameter, top1+top2)

		return top1
	}

	dfs(0)

	return diameter
}

func Max[T cmp.Ordered](nums ...T) T {
	m := nums[0]

	for _, n := range nums {
		if n > m {
			m = n
		}
	}

	return m
}
