package golang

func largestUniqueNumber(nums []int) int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	largest := -1

	for num, times := range cnt {
		if times != 1 {
			continue
		}

		largest = Max(largest, num)
	}

	return largest
}
