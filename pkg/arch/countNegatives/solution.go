package countNegatives

func countNegatives(grid [][]int) int {
	counter := 0
	n := 0

	if len(grid) > 0 {
		n = len(grid[0]) - 1
	}

	for _, raw := range grid {
		for i := n; i >= 0; i-- {
			if raw[i] < 0 {
				counter += 1
			}
		}
	}

	return counter
}
