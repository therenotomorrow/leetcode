package golang

func reverse(x int) int {
	y := 0

	for x != 0 {
		digit := x % 10

		if y > MaxInt32Overflow || y < MinInt32Overflow {
			return 0
		}

		x /= 10
		y = y*10 + digit
	}

	return y
}
