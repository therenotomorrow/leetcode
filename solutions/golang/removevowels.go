package golang

func removeVowels(s string) string {
	runes := make([]rune, 0)

	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			continue
		}

		runes = append(runes, r)
	}

	return string(runes)
}
