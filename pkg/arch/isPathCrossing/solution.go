package isPathCrossing

func isPathCrossing(path string) bool {
	set := make(map[[2]int]bool)
	cur := [2]int{0, 0}

	for _, p := range path {
		set[cur] = true

		switch p {
		case 'N':
			cur[1]++
		case 'S':
			cur[1]--
		case 'E':
			cur[0]++
		case 'W':
			cur[0]--
		}

		if set[cur] {
			return true
		}
	}

	return false
}
