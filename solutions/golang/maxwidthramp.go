package golang

func maxWidthRamp(nums []int) int {
	const previous = 2

	fromRtoLMax := make([]int, len(nums))

	fromRtoLMax[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - previous; i >= 0; i-- {
		fromRtoLMax[i] = Max(fromRtoLMax[i+1], nums[i])
	}

	width := 0

	for left, right := 0, 0; right < len(nums); right++ {
		for left < right && nums[left] > fromRtoLMax[right] {
			left++
		}

		width = Max(width, right-left)
	}

	return width
}
