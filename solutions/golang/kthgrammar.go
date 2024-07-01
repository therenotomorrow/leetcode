package golang

func kthGrammar(_ int, k int) int {
	const half = 2

	numBits := 0

	for k--; k > 0; {
		numBits += k & 1
		k >>= 1
	}

	return numBits % half
}
