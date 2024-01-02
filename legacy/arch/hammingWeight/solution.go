package hammingWeight

func hammingWeight(num uint32) int {
	cnt := 0

	for num > 0 {
		cnt += int(num & 1)
		num >>= 1
	}

	return cnt
}
