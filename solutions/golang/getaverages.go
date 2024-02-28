package golang

func getAverages(nums []int, k int) []int {
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	ans := make([]int, 0)

	for i := range nums {
		var mean int

		if i < k || i+k >= len(nums) {
			mean = -1
		} else {
			mean = (prefix[i+k] - prefix[i-k] + nums[i-k]) / (2*k + 1)
		}

		ans = append(ans, mean)
	}

	return ans
}
