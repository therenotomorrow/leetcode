package golang

func countGood(nums []int, target int) int64 {
	ans := 0
	cnt := make(map[int]int)

	var right, pairs int

	for _, num := range nums {
		for ; pairs < target && right < len(nums); right++ {
			pairs += cnt[nums[right]]
			cnt[nums[right]]++
		}

		if pairs >= target {
			ans += len(nums) - right + 1
		}

		cnt[num]--
		pairs -= cnt[num]
	}

	return int64(ans)
}
