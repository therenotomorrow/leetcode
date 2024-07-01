package golang

func mySqrt(num int) int {
	const half = 2

	if num <= 1 {
		return num
	}

	left := 0
	right := num / half

	for left <= right {
		mid := (right + left) / half

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
