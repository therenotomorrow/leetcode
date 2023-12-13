package solutions

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (right + left) / 2

		switch curr := nums[mid]; {
		case curr < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		default:
			return mid
		}
	}

	return -1
}
