package golang

func isPerfectSquare(num int) bool {
	// also Newton's method could be used there
	if num <= 1 {
		return true
	}

	left := 0
	right := num / Half

	for left <= right {
		mid := (right + left) / Half

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
