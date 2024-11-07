package golang

func largestCombination(candidates []int) int {
	const maxBits = 24

	largest := 0

	for bits := range maxBits {
		cnt := 0

		for _, candidate := range candidates {
			if candidate&(1<<bits) > 0 {
				cnt++
			}
		}

		largest = Max(largest, cnt)
	}

	return largest
}
