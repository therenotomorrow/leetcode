package golang

func sumHappy(n int) int {
	newN := 0

	for n > 0 {
		d := n % Digits
		newN += d * d
		n /= Digits
	}

	return newN
}

func isHappy(n int) bool {
	// find cycle, could use fast and slow runners
	seen := NewSet[int]()

	for n > 1 {
		n = sumHappy(n)

		if seen.Contains(n) {
			return false
		}

		seen.Add(n)
	}

	return n == 1
}
