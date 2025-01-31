package golang

func eventualSafeNodes(graph [][]int) []int {
	var (
		dfs     func(curr int) bool
		visited = make([]bool, len(graph))
		unsafe  = make([]bool, len(graph))
	)

	dfs = func(curr int) bool {
		if unsafe[curr] {
			return true
		}

		if visited[curr] {
			return false
		}

		visited[curr] = true
		unsafe[curr] = true

		for _, neighbor := range graph[curr] {
			if dfs(neighbor) {
				return true
			}
		}

		unsafe[curr] = false

		return false
	}

	for curr := range graph {
		_ = dfs(curr)
	}

	safe := make([]int, 0)

	for curr := range graph {
		if !unsafe[curr] {
			safe = append(safe, curr)
		}
	}

	return safe
}
