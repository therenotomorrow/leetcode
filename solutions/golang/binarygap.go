package golang

func binaryGap(n int) int {
	bin := make([]int, 0)

	for ; n > 0; n /= Base2 {
		bin = append(bin, n%Base2)
	}

	start := 0

	for i, digit := range bin {
		if digit == 1 {
			start = i

			break
		}
	}

	gap := 0

	for end := start; end < len(bin); end++ {
		if bin[end] == 1 {
			gap = Max(gap, end-start)
			start = end
		}
	}

	return gap
}
