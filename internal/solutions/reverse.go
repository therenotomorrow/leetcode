package solutions

import "github.com/therenotomorrow/leetcode/internal/structs"

func reverse(x int) int {
	y := 0

	for x != 0 {
		digit := x % 10

		if y > structs.MaxInt32Overflow || y < structs.MinInt32Overflow {
			return 0
		}

		x /= 10
		y = y*10 + digit
	}

	return y
}
