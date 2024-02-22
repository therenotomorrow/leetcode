package golang

func destCity(paths [][]string) string {
	set := make(map[string]bool)

	for _, path := range paths {
		src := path[0]
		set[src] = true
	}

	for _, path := range paths {
		dst := path[1]

		if !set[dst] {
			return dst
		}
	}

	// impossible via task description
	return ""
}
