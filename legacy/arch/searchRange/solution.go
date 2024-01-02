package searchRange

func searchRange(nums []int, target int) []int {
	i, ok := binarySearch(nums, target)
	if !ok {
		return []int{-1, -1}
	}

	firstPos := i
	for i := i; i >= 0 && nums[i] == target; i-- {
		firstPos = i
	}

	lastPos := i
	for i := i; i < len(nums) && nums[i] == target; i++ {
		lastPos = i
	}

	return []int{firstPos, lastPos}
}

func binarySearch(nums []int, target int) (int, bool) {
	var (
		l = 0
		r = len(nums) - 1
	)

	for l <= r {
		m := (l + r) / 2

		switch {
		case target == nums[m]:
			return m, true
		case target < nums[m]:
			r = m - 1
		default:
			l = m + 1
		}
	}

	return 0, false
}
