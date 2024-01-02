package longestArithSeqLength

func maxi(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestArithSeqLength(nums []int) int {
	visited := make([]map[int]int, len(nums))
	res := 0

	for right := 0; right < len(nums); right++ {
		visited[right] = make(map[int]int)

		for left := 0; left < right; left++ {
			diff := nums[right] - nums[left]

			if val, ok := visited[left][diff]; !ok {
				visited[right][diff] = 2
			} else {
				visited[right][diff] = val + 1
			}

			res = maxi(res, visited[right][diff])
		}
	}

	return res
}
