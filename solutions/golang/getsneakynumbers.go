package golang

func getSneakyNumbers(nums []int) []int {
	const size = 2

	var (
		ans = make([]int, 0, size)
		cnt = make(map[int]int)
	)

	for _, num := range nums {
		cnt[num]++

		if cnt[num] > 1 {
			ans = append(ans, num)
		}
	}

	return ans
}
