package golang

func mySqrt(num int) int {
	if num <= 1 {
		return num
	}

	left := 0
	right := num / Half

	for left <= right {
		mid := (right + left) / Half

		if mid*mid == num {
			return mid
		}

		if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}
