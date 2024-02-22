package golang

func maxScore(s string) int {
	ones := 0

	for _, r := range s {
		if r == '1' {
			ones++
		}
	}

	zeros := 0
	score := 0

	for _, r := range s[:len(s)-1] {
		switch r {
		case '0':
			zeros++
		case '1':
			ones--
		}

		score = Max(score, zeros+ones)
	}

	return score
}
