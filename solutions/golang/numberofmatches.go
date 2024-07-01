package golang

func numberOfMatches(n int) int {
	const half = 2

	// with simulation
	matches := 0

	for n > 1 {
		matches += (n - n%half) / half
		n = (n + n%half) / half
	}

	// with logic: because the winner will only one - that's why `n - 1`
	return matches
}
