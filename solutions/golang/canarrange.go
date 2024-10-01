package golang

func canArrange(arr []int, k int) bool {
	cnt := make(map[int]int)
	for _, num := range arr {
		cnt[(num%k+k)%k]++
	}

	for remainder, times := range cnt {
		if remainder == 0 && times%2 != 0 || remainder != 0 && cnt[remainder] != cnt[k-remainder] {
			return false
		}
	}

	return true
}
