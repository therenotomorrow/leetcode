package golang

func isPathCrossing(path string) bool {
	set := NewSet[[2]int]()
	cur := [2]int{0, 0}

	for _, p := range path {
		set.Add(cur)

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

		if set.Contains(cur) {
			return true
		}
	}

	return false
}
