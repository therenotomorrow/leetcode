package golang

// var solution = func(read4 func([]byte) int) func([]byte, int) int.
func solution(read4 func([]byte) int) func([]byte, int) int {
	return func(buf []byte, n int) int {
		pnt := 0
		buf4 := make([]byte, 4)

		for pnt < n {
			// read4 will override exist buffer and signal about how much read
			read := read4(buf4)

			if read == 0 {
				break
			}

			if pnt+read > n {
				read = n - pnt
			}

			for j := 0; j < read; j++ {
				buf[pnt] = buf4[j]
				pnt++
			}
		}

		return pnt
	}
}
