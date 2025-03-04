package golang

func checkPowersOfThree(n int) bool {
	const two = 2

	for n != 1 {
		if n%3 == two {
			return false
		}

		n /= 3
	}

	return true
}
