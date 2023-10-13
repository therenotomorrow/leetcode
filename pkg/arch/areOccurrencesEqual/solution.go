package areOccurrencesEqual

func areOccurrencesEqual(s string) bool {
	freq := make(map[rune]int)
	for _, r := range s {
		freq[r]++
	}

	set := make(map[int]struct{})
	for _, v := range freq {
		set[v] = struct{}{}
	}

	return len(set) == 1
}
