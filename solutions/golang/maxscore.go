package golang

func maxScore(s string) int {
	var ones, zeros, score int

	for _, r := range s {
		if r == '1' {
			ones++
		}
	}

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
