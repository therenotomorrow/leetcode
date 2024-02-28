package golang

func calculateTime(keyboard string, word string) int {
	keys := make(map[rune]int, 26)

	for i, key := range keyboard {
		keys[key] = i
	}

	var time, start int

	for _, r := range word {
		end := keys[r]
		time += Abs(end - start)
		start = end
	}

	return time
}
