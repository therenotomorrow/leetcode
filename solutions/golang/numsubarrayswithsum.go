package golang

func numSubarraysWithSum(nums []int, goal int) int {
	sums := make(map[int]int)
	sums[0] = 1

	var ans, curr int
	for _, num := range nums {
		curr += num
		ans += sums[curr-goal]
		sums[curr]++
	}

	return ans
}
