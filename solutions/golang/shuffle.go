package golang

func shuffle(nums []int, n int) []int {
	ans := make([]int, Double*n)

	for i, j := 0, 0; i < n; i++ {
		ans[j] = nums[i]
		ans[j+1] = nums[n+i]

		j += Double
	}

	return ans
}
