package checkIfPangram

func checkIfPangram(sentence string) bool {
	keys := make(map[rune]bool)

	for _, r := range sentence {
		keys[r] = true
	}

	return len(keys) == 26
}
