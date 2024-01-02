package minimumOneBitOperations

func minimumOneBitOperations(n int) int {
	if n == 0 {
		return 0
	}

	k := 0
	c := 1

	for c*2 <= n {
		c *= 2
		k++
	}

	pow := 1
	for i := 0; i < k+1; i++ {
		pow *= 2
	}

	return pow - minimumOneBitOperations(n^c) - 1
}
