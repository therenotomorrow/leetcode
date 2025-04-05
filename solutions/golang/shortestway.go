package golang

func shortestWay(source string, target string) int {
	set := NewSet[rune]()
	for _, r := range source {
		set.Add(r)
	}

	for _, r := range target {
		if !set.Contains(r) {
			return -1
		}
	}

	i := 0
	ways := 0
	size := len(source)

	for _, r := range target {
		if i == 0 {
			ways++
		}

		for rune(source[i]) != r {
			i = (i + 1) % size

			if i == 0 {
				ways++
			}
		}

		i = (i + 1) % size
	}

	return ways
}
