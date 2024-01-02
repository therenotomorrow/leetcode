package reverseBits

func reverseBits(num uint32) uint32 {
	var (
		ans   uint32
		shift = 31
	)

	for num != 0 {
		ans += (num & 1) << shift
		num >>= 1
		shift--
	}

	return ans
}
