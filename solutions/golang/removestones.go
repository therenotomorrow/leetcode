package golang

func removeStones(stones [][]int) int {
	var (
		graph   = make([][]int, len(stones))
		visited = make([]bool, len(stones))
		dfs     func(curr int)
	)

	dfs = func(curr int) {
		visited[curr] = true

		for _, neighbor := range graph[curr] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	for i := range stones {
		for j := i + 1; j < len(stones); j++ {
			if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}

	connected := 0

	for i := range stones {
		if !visited[i] {
			dfs(i)

			connected++
		}
	}

	return len(stones) - connected
}
