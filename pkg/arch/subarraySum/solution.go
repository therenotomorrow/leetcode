package subarraySum

func subarraySum(nums []int, k int) int {
	counts := make(map[int]int)
	counts[0] = 1

	var ans, curr int
	for _, num := range nums {
		curr += num
		ans += counts[curr-k]
		counts[curr]++
	}

	return ans
}
