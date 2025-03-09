package golang

func numberOfAlternatingGroups(colors []int, k int) int {
	groups := 0
	currSize := 1
	currColor := colors[0]

	for i := 1; i < len(colors)+k-1; i++ {
		j := i % len(colors)

		if currColor == colors[j] {
			currColor = colors[j]
			currSize = 1

			continue
		}

		currSize++

		if currSize >= k {
			groups++
		}

		currColor = colors[j]
	}

	return groups
}
