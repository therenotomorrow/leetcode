package golang

func validWordSquare(words []string) bool {
	for i := range len(words) {
		for j := range len(words[i]) {
			// index error restrictions
			if j >= len(words) || i >= len(words[j]) {
				return false
			}

			// task description restrictions
			if words[i][j] != words[j][i] {
				return false
			}
		}
	}

	return true
}
