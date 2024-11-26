package findChampion

func findChampion(n int, edges [][]int) int {
	indegree := make([]int, n)

	for _, edge := range edges {
		indegree[edge[1]]++
	}

	champ := -1
	champCount := 0

	for i, val := range indegree {
		if val == 0 {
			champCount++
			champ = i
		}
	}

	if champCount > 1 {
		return -1
	}

	return champ
}
