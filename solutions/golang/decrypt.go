package golang

func decrypt(code []int, k int) []int {
	decrypted := make([]int, len(code))
	step := 1

	if k < 0 {
		step = -step
		k = -k
	}

	for i := range code {
		window := k

		for j := i + step; window > 0; j += step {
			switch {
			case j >= len(code):
				j = 0
			case j < 0:
				j = len(code) - 1
			}

			decrypted[i] += code[j]

			window--
		}
	}

	return decrypted
}
