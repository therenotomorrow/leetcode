package golang

func passThePillow(n int, time int) int {
	full := time / (n - 1)
	rest := time % (n - 1)

	if full%2 == 0 {
		return rest + 1
	}

	return n - rest
}
