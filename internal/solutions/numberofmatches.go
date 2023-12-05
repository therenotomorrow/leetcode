package solutions

func numberOfMatches(n int) int {
	// with simulation
	matches := 0

	for n > 1 {
		matches += (n - n%2) / 2
		n = (n + n%2) / 2
	}

	// with logic: because the winner will only one - that's why `n - 1`
	return matches
}
