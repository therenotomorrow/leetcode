package golang

func uniqueOccurrences(arr []int) bool {
	occ := make(map[int]int)
	uni := make(map[int]struct{})

	for _, num := range arr {
		occ[num]++
	}

	for _, times := range occ {
		_, ok := uni[times]
		if ok {
			return false
		}

		uni[times] = struct{}{}
	}

	return true
}
