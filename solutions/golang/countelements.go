package golang

func countElements(arr []int) int {
	cnt := 0
	set := NewSet[int]()

	for _, num := range arr {
		set.Add(num)
	}

	for _, num := range arr {
		if set.Contains(num + 1) {
			cnt++
		}
	}

	return cnt
}
