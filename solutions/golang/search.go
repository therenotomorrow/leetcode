package golang

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (right + left) / Half

		switch curr := nums[mid]; {
		case curr < target:
			left = mid + 1
		case curr > target:
			right = mid - 1
		default:
			return mid
		}
	}

	return -1
}
