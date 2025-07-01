package golang

func findLHS(nums []int) int {
	ans := 0
	cnt := make(map[int]int)

	for _, num := range nums {
		cnt[num]++

		if val, ok := cnt[num-1]; ok {
			ans = Max(ans, cnt[num]+val)
		}

		if val, ok := cnt[num+1]; ok {
			ans = Max(ans, cnt[num]+val)
		}
	}

	return ans
}
