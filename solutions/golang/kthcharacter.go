package golang

func kthCharacter(k int) byte {
	shifts := 0

	for k > 1 {
		power := 1
		for power<<1 < k {
			power <<= 1
		}

		if k >= power {
			k -= power
			shifts++
		}
	}

	return byte(int('a') + shifts)
}
