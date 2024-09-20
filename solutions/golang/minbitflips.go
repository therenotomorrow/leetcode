package golang

func minBitFlips(start int, goal int) int {
	cnt := 0

	for rest := start ^ goal; rest > 0; rest >>= 1 {
		if rest%2 == 1 {
			cnt++
		}
	}

	return cnt
}
