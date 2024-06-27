package golang

func findCenter(edges [][]int) int {
	cnt := make(map[int]int)

	for _, edge := range edges {
		cnt[edge[0]]++
		cnt[edge[1]]++
	}

	center := 0
	maxCnt := 0
	for node, times := range cnt {
		if times > maxCnt {
			maxCnt = times
			center = node
		}
	}

	return center
}
