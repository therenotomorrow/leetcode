package hammingDistance

func hammingDistance(x int, y int) int {
	cnt := 0

	if x > y {
		for x != 0 {
			a := x & 1
			b := y & 1

			if a != b {
				cnt++
			}

			x >>= 1
			y >>= 1
		}
	} else {
		for y != 0 {
			a := x & 1
			b := y & 1

			if a != b {
				cnt++
			}

			x >>= 1
			y >>= 1
		}
	}

	return cnt
}
