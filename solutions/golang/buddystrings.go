package golang

func buddyStrings(text string, goal string) bool { //nolint:cyclop
	if len(text) != len(goal) {
		return false
	}

	if text == goal {
		cnt := make(map[rune]int)

		for _, runa := range text {
			cnt[runa]++

			if cnt[runa] > 1 {
				return true
			}
		}

		return false
	}

	one := -1
	two := -1

	for i := range text {
		if text[i] == goal[i] {
			continue
		}

		switch {
		case one == -1:
			one = i
		case two == -1:
			two = i
		default:
			return false
		}
	}

	return two != -1 && text[one] == goal[two] && text[two] == goal[one]
}
