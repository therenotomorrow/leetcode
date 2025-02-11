package golang

func removeOccurrences(text string, part string) string {
	runes := make([]rune, 0)
	size := len(part)

	for _, runa := range text {
		runes = append(runes, runa)
		diff := len(runes) - size

		if diff >= 0 && string(runes[diff:]) == part {
			runes = runes[:diff]
		}
	}

	return string(runes)
}
