package guessNumber

func guess(num int) int {
	pick := 6

	if num > pick {
		return -1
	}

	if num < pick {
		return 1
	}

	return 0
}

func guessNumber(n int) int {
	l := 1
	r := n

	for l < r {
		n = (r + l) / 2
		g := guess(n)

		switch g {
		case -1:
			r = n - 1
		case 1:
			l = n + 1
		default:
			return n
		}
	}

	return r
}
