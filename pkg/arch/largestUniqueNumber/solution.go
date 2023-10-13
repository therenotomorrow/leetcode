package largestUniqueNumber

func largestUniqueNumber(nums []int) int {
	visited := make(map[int]int)

	for _, num := range nums {
		visited[num]++
	}

	max := -1
	for num, cnt := range visited {
		if cnt != 1 {
			continue
		}

		if num > max {
			max = num
		}
	}

	return max
}
