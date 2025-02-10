package golang

func countBadPairs(nums []int) int64 {
	beds := 0
	cnt := make(map[int]int)

	for i, num := range nums {
		diff := i - num
		goods := cnt[diff]
		beds += i - goods
		cnt[diff] = goods + 1
	}

	return int64(beds)
}
