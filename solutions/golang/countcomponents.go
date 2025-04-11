package golang

func countComponents(n int, edges [][]int) int {
	var (
		ans     int
		dfs     func(curr int)
		nodes   = make(map[int][]int)
		visited = make([]bool, n)
	)

	for _, edge := range edges {
		nodes[edge[0]] = append(nodes[edge[0]], edge[1])
		nodes[edge[1]] = append(nodes[edge[1]], edge[0])
	}

	dfs = func(curr int) {
		visited[curr] = true

		for _, node := range nodes[curr] {
			if visited[node] {
				continue
			}

			dfs(node)
		}
	}

	for i := range n {
		if visited[i] {
			continue
		}

		ans++

		dfs(i)
	}

	return ans
}
