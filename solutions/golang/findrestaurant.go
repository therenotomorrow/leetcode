package golang

func findRestaurant(list1 []string, list2 []string) []string {
	words := make(map[string]int)
	for _, word := range list1 {
		words[word] = -1
	}

	for i, word := range list2 {
		if _, ok := words[word]; ok {
			words[word] += i + 1
		}
	}

	ans := make([]string, 0)
	minVal := len(list1) + len(list2)

	for i, word := range list1 {
		j := words[word]
		if j == -1 {
			delete(words, word)

			continue
		}

		switch newVal := i + j; {
		case newVal < minVal:
			minVal = newVal
			ans = []string{word}
		case newVal == minVal:
			ans = append(ans, word)
		}
	}

	return ans
}
