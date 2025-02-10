package golang

func queryResults(_ int, queries [][]int) []int {
	balls := make(map[int]int)
	colors := make(map[int]int)
	ans := make([]int, len(queries))

	for i := range queries {
		ball := queries[i][0]

		if lastColor := balls[ball]; lastColor != 0 {
			colors[lastColor]--

			if colors[lastColor] == 0 {
				delete(colors, lastColor)
			}
		}

		color := queries[i][1]
		balls[ball] = color
		colors[color]++

		ans[i] = len(colors)
	}

	return ans
}
