package solutions

func hammingWeight(num uint32) int {
	cnt := uint32(0)

	for ; num > 0; num >>= 1 {
		cnt += num & 1
	}

	return int(cnt)
}
