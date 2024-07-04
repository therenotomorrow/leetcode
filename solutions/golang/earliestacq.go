package golang

import "sort"

func earliestAcqConnected(graph [][]int, n int) bool {
	var (
		visited = make([]bool, n)
		dfs     func(int)
	)

	dfs = func(node int) {
		visited[node] = true

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	dfs(0)

	for i := range n {
		if !visited[i] {
			return false
		}
	}

	return true
}

func earliestAcq(logs [][]int, n int) int {
	sort.SliceStable(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	graph := make([][]int, n)
	for i := range n {
		graph[i] = make([]int, 0)
	}

	for _, log := range logs {
		friend1 := log[1]
		friend2 := log[2]

		graph[friend1] = append(graph[friend1], friend2)
		graph[friend2] = append(graph[friend2], friend1)

		if earliestAcqConnected(graph, n) {
			return log[0]
		}
	}

	return -1
}
