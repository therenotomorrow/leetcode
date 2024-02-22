package golang

func shuffle(nums []int, n int) []int {
	ans := make([]int, 2*n)

	for i, j := 0, 0; i < n; i++ {
		ans[j] = nums[i]
		ans[j+1] = nums[n+i]

		j += 2
	}

	return ans
}
