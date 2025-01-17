package golang

func doesValidArrayExist(derived []int) bool {
	xor := 0

	for _, num := range derived {
		xor ^= num
	}

	return xor == 0
}
