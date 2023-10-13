package maxNumberOfBalloons

func maxNumberOfBalloons(text string) int {
	letters := make(map[rune]int)

	for _, r := range text {
		switch r {
		case 'b', 'a', 'l', 'o', 'n':
			letters[r]++
		}
	}

	// special case
	if len(letters) != 5 {
		return 0
	}

	min := len(text)
	for r, cnt := range letters {
		switch r {
		case 'l', 'o':
			cnt /= 2
		}

		if cnt < min {
			min = cnt
		}
	}

	return min
}
