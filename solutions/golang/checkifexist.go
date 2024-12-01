package golang

func checkIfExist(arr []int) bool {
	cnt := make(map[int]int)

	for i, num := range arr {
		cnt[num] = i
	}

	for j, num := range arr {
		i, ok := cnt[2*num]
		if ok && i != j {
			return true
		}
	}

	return false
}
