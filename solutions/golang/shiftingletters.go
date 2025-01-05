package golang

func shiftingLetters(text string, shifts [][]int) string {
	prefix := make([]int, len(text)+1)
	shifter := func(char rune, diff int) rune {
		add := (char - 'a' + rune(diff)) % Alphabet

		return 'a' + (Alphabet+add)%Alphabet
	}

	for _, shift := range shifts {
		if shift[2] == 0 {
			shift[2] = -1
		}

		prefix[shift[0]] += shift[2]
		prefix[shift[1]+1] -= shift[2]
	}

	for i := 1; i < len(prefix); i++ {
		prefix[i] += prefix[i-1]
	}

	runes := []rune(text)
	for i, char := range runes {
		runes[i] = shifter(char, prefix[i])
	}

	return string(runes)
}
