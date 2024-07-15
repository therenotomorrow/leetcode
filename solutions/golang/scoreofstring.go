package golang

func scoreOfString(s string) int {
	score := 0

	for i := range len(s) - 1 {
		score += Abs(int(s[i+1]) - int(s[i]))
	}

	return score
}
