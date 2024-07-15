package golang

func orArray(nums []int) []int {
	ans := make([]int, len(nums)-1)

	for i := range len(nums) - 1 {
		ans[i] = nums[i] | nums[i+1]
	}

	return ans
}
