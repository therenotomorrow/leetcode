package golang

func numberOfPairs(nums []int) []int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	var yes, not int

	for _, times := range cnt {
		if times%2 == 1 {
			not++
		}

		yes += times / Half
	}

	return []int{yes, not}
}
