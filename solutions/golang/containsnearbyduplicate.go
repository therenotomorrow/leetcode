package golang

func containsNearbyDuplicate(nums []int, k int) bool {
	cnt := make(map[int]int)

	for i, num := range nums {
		j, ok := cnt[num]

		if ok && Abs(i-j) <= k {
			return true
		}

		cnt[num] = i
	}

	return false
}
