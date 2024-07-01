package golang

func isPerfectSquare(num int) bool {
	const half = 2

	// also Newton's method could be used there
	if num <= 1 {
		return true
	}

	left := 0
	right := num / half

	for left <= right {
		mid := (right + left) / half

		if mid*mid == num {
			return true
		}

		if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
