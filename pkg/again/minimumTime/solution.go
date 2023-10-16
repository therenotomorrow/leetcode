package minimumTime

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimumTime(n int, relations [][]int, time []int) int {
	// attention: try to find other clean way then editorial implementation later
	var (
		cache   = make(map[int]int)
		graph   = make(map[int][]int, len(relations))
		dynamic func(int) int
	)

	dynamic = func(node int) int {
		if graph[node] == nil {
			return time[node]
		}

		if val, ok := cache[node]; ok {
			return val
		}

		ans := 0
		for _, neighbor := range graph[node] {
			ans = max(ans, dynamic(neighbor))
		}

		cache[node] = time[node] + ans

		return time[node] + ans
	}

	for _, relation := range relations {
		key := relation[0] - 1
		val := relation[1] - 1

		if graph[key] == nil {
			graph[key] = make([]int, 0)
		}

		graph[key] = append(graph[key], val)
	}

	ans := 0
	for node := 0; node < n; node++ {
		ans = max(ans, dynamic(node))
	}

	return ans
}
