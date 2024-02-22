package golang

func evenOddBit(n int) []int {
	ans := []int{0, 0}

	for isEven := false; n > 0; n >>= 1 {
		isEven = !isEven

		if n&1 != 1 {
			continue
		}

		if isEven {
			ans[0]++
		} else {
			ans[1]++
		}
	}

	return ans
}
