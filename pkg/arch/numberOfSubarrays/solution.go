package numberOfSubarrays

func numberOfSubarrays(nums []int, k int) int {
	counts := make(map[int]int)
	counts[0] = 1

	var ans, curr int
	for _, num := range nums {
		curr += num % 2
		ans += counts[curr-k]
		counts[curr]++
	}

	return ans
}
