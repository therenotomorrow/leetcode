package golang

func removeVowels(str string) string {
	runes := make([]rune, 0)

	for _, runa := range str {
		switch runa {
		case 'a', 'e', 'i', 'o', 'u':
			continue
		}

		runes = append(runes, runa)
	}

	return string(runes)
}
