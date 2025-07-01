package golang

func maxAdjacentDistance(nums []int) int {
	ans := 0

	nums = append(nums, nums[0])

	for i := range len(nums) - 1 {
		ans = Max(ans, Abs(nums[i]-nums[i+1]))
	}

	return ans
}
