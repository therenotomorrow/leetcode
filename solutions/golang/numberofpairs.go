package golang

func numberOfPairs(nums []int) []int {
	const half = 2

	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	var yes, not int

	for _, times := range cnt {
		if times%2 == 1 {
			not++
		}

		yes += times / half
	}

	return []int{yes, not}
}
