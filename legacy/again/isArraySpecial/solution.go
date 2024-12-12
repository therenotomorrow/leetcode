package isArraySpecial

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	// Precompute prefix array where prefix[i] == 1 if nums[i] and nums[i+1] have the same parity
	prefix := make([]int, n)
	for i := 0; i < n-1; i++ {
		if nums[i]%2 == nums[i+1]%2 {
			prefix[i+1] = prefix[i] + 1
		} else {
			prefix[i+1] = prefix[i]
		}
	}

	result := make([]bool, len(queries))
	for i, query := range queries {
		from, to := query[0], query[1]
		result[i] = prefix[to]-prefix[from] == 0
	}

	return result
}
