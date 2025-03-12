package golang

func maximumCount(nums []int) int {
	const half = 2

	binarySearch := func(nums []int, cmp func(curr int) bool) int {
		left := 0
		right := len(nums) - 1

		for left <= right {
			mid := (right + left) / half

			switch curr := nums[mid]; {
			case cmp(curr):
				left = mid + 1
			default:
				right = mid - 1
			}
		}

		return left
	}

	neg := binarySearch(nums, func(curr int) bool { return curr < 0 })
	pos := binarySearch(nums, func(curr int) bool { return curr <= 0 })

	return Max(neg, len(nums)-pos)
}
