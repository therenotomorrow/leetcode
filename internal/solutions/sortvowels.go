package solutions

func sortVowels(s string) string {
	runes := []rune(s)
	count := map[rune]int{'A': 0, 'E': 0, 'I': 0, 'O': 0, 'U': 0, 'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}

	for i, r := range runes {
		_, ok := count[r]
		if ok {
			count[r]++

			runes[i] = '_'
		}
	}

	point := 0
	order := []rune("AEIOUaeiou")

	for i, r := range runes {
		if r != '_' {
			continue
		}

		for count[order[point]] == 0 {
			point++
		}

		runes[i] = order[point]
		count[order[point]]--
	}

	return string(runes)
}
