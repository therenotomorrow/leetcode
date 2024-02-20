package solutions

func isCircularSentence(sentence string) bool {
	for i, r := range sentence {
		if r != ' ' {
			continue
		}

		if sentence[i-1] != sentence[i+1] {
			return false
		}
	}

	return sentence[0] == sentence[len(sentence)-1]
}
