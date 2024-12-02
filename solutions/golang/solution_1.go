package golang

func solution1(read4 func([]byte) int) func([]byte, int) int {
	const bufSize = 4

	return func(buf []byte, n int) int {
		pnt := 0
		buf4 := make([]byte, bufSize)

		for pnt < n {
			// read4 will override exist buffer and signal about how much read
			read := read4(buf4)

			if read == 0 {
				break
			}

			if pnt+read > n {
				read = n - pnt
			}

			for j := range read {
				buf[pnt] = buf4[j]
				pnt++
			}
		}

		return pnt
	}
}
