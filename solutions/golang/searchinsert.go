package golang

func searchInsert(nums []int, target int) int {
	const half = 2

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (right + left) / half

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
