package uniqueOccurrences

func uniqueOccurrences(arr []int) bool {
	occ := make(map[int]int)
	uni := make(map[int]struct{})

	for _, num := range arr {
		occ[num]++
	}

	for _, v := range occ {
		_, ok := uni[v]
		if ok {
			return false
		}
		uni[v] = struct{}{}
	}

	return true
}
