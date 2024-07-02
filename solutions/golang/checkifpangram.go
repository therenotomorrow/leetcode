package golang

func checkIfPangram(sentence string) bool {
	set := NewSet[rune]()

	for _, r := range sentence {
		set.Add(r)

		if set.Size() == Alphabet {
			return true
		}
	}

	return false
}
