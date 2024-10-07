package golang

func findKthNumber(limit int, k int) int {
	stepper := func(num int) int {
		steps := 0
		curr, next := num, num+1

		for curr <= limit {
			steps += Min(next, limit+1) - curr
			curr *= Digits
			next *= Digits
		}

		return steps
	}

	num := 1

	for k--; k > 0; {
		steps := stepper(num)

		if k >= steps {
			k -= steps
			num++
		} else {
			k--
			num *= Digits
		}
	}

	return num
}
