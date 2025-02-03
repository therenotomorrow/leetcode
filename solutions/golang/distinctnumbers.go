package golang

func distinctNumbers(nums []int, k int) []int {
	i := 0

	ans := make([]int, 0)
	cnt := make(map[int]int)

	for j, num := range nums {
		if j-i < k {
			cnt[num]++

			continue
		}

		ans = append(ans, len(cnt))

		cnt[nums[i]]--
		if cnt[nums[i]] == 0 {
			delete(cnt, nums[i])
		}

		cnt[num]++
		i++
	}

	return append(ans, len(cnt))
}
