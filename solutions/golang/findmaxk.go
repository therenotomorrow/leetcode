package golang

func findMaxK(nums []int) int {
	curr := -1
	cnt := make(map[int]int)

	for _, num := range nums {
		cnt[num] = num
	}

	for _, num := range nums {
		_, ok := cnt[-num]
		if !ok {
			continue
		}

		curr = Max(curr, num)
	}

	return curr
}
