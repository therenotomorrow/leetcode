package pivotIndex

func pivotIndex(nums []int) int {
	l, r := 0, 0

	for i := 0; i < len(nums); i++ {
		r += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if l == r-l-nums[i] {
			return i
		}
		l += nums[i]
	}

	return -1
}
