package golang

func maxScore(str string) int {
	var ones, zeros, score int

	for _, r := range str {
		if r == '1' {
			ones++
		}
	}

	for _, r := range str[:len(str)-1] {
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
