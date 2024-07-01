package golang

func shuffle(nums []int, n int) []int {
	const double = 2

	ans := make([]int, double*n)

	for i, j := 0, 0; i < n; i++ {
		ans[j] = nums[i]
		ans[j+1] = nums[n+i]

		j += double
	}

	return ans
}
