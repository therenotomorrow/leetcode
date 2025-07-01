package golang

func maximumDifference(nums []int) int {
	ans := -1
	lastMin := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] <= lastMin {
			lastMin = nums[i]
		} else {
			ans = Max(ans, nums[i]-lastMin)
		}
	}

	return ans
}
