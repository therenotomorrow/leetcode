package numIdenticalPairs

func numIdenticalPairs(nums []int) int {
	m := make(map[int]int)

	var cnt int
	for _, num := range nums {
		cnt += m[num]
		m[num]++
	}

	return cnt
}
