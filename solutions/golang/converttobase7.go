package golang

import "slices"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	sign := false
	if num < 0 {
		sign = true
		num = -num
	}

	base7 := make([]rune, 0)

	for num > 0 {
		base7 = append(base7, rune(num%Base7)+'0')
		num /= Base7
	}

	slices.Reverse(base7)

	if sign {
		return "-" + string(base7)
	}

	return string(base7)
}
