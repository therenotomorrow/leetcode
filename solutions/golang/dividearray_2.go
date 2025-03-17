package golang

func divideArray2(nums []int) bool {
	cnt := make(map[int]int)

	for _, num := range nums {
		cnt[num]++
	}

	for _, times := range cnt {
		if times%2 != 0 {
			return false
		}
	}

	return true
}
