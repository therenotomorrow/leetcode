package prefixSum

func answerQueries(nums []int, queries [][]int, limit int) []bool {
	if len(nums) < 1 {
		return []bool{}
	}

	prefixes := make([]int, len(nums))

	prefixes[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixes[i] = prefixes[i-1] + nums[i]
	}

	result := make([]bool, len(queries))

	for i, query := range queries {
		sum := prefixes[query[1]] - prefixes[query[0]] + nums[query[0]]
		result[i] = sum < limit
	}

	return result
}
