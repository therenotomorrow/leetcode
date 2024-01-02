package solutions

func kthGrammar(_ int, k int) int {
	numBits := 0

	for k--; k > 0; {
		numBits += k & 1
		k >>= 1
	}

	return numBits % 2
}
