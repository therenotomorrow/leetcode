package checkingForExistence

func findNumbers(nums []int) []int {
	set := make(map[int]int)
	for _, num := range nums {
		set[num] = num
	}

	ans := make([]int, 0)
	for _, num := range nums {
		_, ok1 := set[num+1]
		_, ok2 := set[num-1]

		if !ok1 && !ok2 {
			ans = append(ans, num)
		}
	}

	return ans
}
