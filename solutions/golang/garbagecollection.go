package golang

func garbageCollection(garbage []string, travel []int) int {
	trucks := map[rune][]int{
		'M': {0, 0},
		'P': {0, 0},
		'G': {0, 0},
	}
	prefixSums := make([]int, len(travel)+1)

	for i, house := range garbage {
		if i == 0 {
			prefixSums[0] = 0
		} else {
			prefixSums[i] = prefixSums[i-1] + travel[i-1]
		}

		for _, g := range house {
			trucks[g][1]++
			trucks[g][0] = prefixSums[i]
		}
	}

	minToPick := 0

	for _, truck := range trucks {
		minToPick += truck[0] + truck[1]
	}

	return minToPick
}
