package golang

func minimumSize(nums []int, maxOperations int) int {
	const half = 2

	divide := func(capacity int) bool {
		operations := 0

		for _, num := range nums {
			operations += (num+capacity-1)/capacity - 1

			if operations > maxOperations {
				return false
			}
		}

		return true
	}

	left := 1
	right := Max(nums...)

	for left < right {
		mid := left + (right-left)/half

		if divide(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
