package findLucky

func findLucky(arr []int) int {
	m := make(map[int]int)

	for _, num := range arr {
		m[num]++
	}

	max := -1
	for num, cnt := range m {
		if num == cnt && num > max {
			max = num
		}
	}

	return max
}
