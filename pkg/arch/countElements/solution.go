package countElements

func countElements(arr []int) int {
	set := make(map[int]bool)
	for _, num := range arr {
		set[num] = true
	}

	cnt := 0
	for _, num := range arr {
		if set[num+1] {
			cnt++
		}
	}

	return cnt
}
