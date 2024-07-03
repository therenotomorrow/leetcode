package golang

func isAlienSorted(words []string, order string) bool {
	ord := make(map[byte]int)

	for i, let := range order {
		ord[byte(let)] = i
	}

	for i := range len(words) - 1 {
		for j := range len(words[i]) {
			if j >= len(words[i+1]) {
				return false
			}

			if words[i][j] != words[i+1][j] {
				if ord[words[i][j]] > ord[words[i+1][j]] {
					return false
				}

				break
			}
		}
	}

	return true
}
