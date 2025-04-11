package golang

func numberOfMatches(n int) int {
	// with simulation
	matches := 0

	for n > 1 {
		matches += (n - n%Half) / Half
		n = (n + n%Half) / Half
	}

	// with logic: because the winner will only one - that's why `n - 1`
	return matches
}
