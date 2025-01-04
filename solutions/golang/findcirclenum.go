package golang

func findCircleNum(isConnected [][]int) int {
	var (
		dfs       func(node int)
		provinces = 0
		visited   = make([]bool, len(isConnected))
	)

	dfs = func(node int) {
		if visited[node] {
			return
		}

		visited[node] = true

		for neighbor, connected := range isConnected[node] {
			if connected == 1 {
				dfs(neighbor)
			}
		}
	}

	for node := range isConnected {
		if !visited[node] {
			provinces++

			dfs(node)
		}
	}

	return provinces
}
