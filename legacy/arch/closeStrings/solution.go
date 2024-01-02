package closeStrings

import "sort"

func closeStrings(word1 string, word2 string) bool {
	// step 0: check simple condition
	if len(word1) != len(word2) {
		return false
	}

	// step 1: check that all characters the same
	runes1 := make(map[rune]int)
	for _, r := range word1 {
		runes1[r]++
	}

	runes2 := make(map[rune]int)
	for _, r := range word2 {
		runes2[r]++

		if _, ok := runes1[r]; !ok {
			return false
		}
	}

	// step 2: check that all frequencies the same
	freq1 := make([]int, 0, len(runes1))
	for _, v := range runes1 {
		freq1 = append(freq1, v)
	}
	sort.Ints(freq1)

	freq2 := make([]int, 0, len(runes2))
	for _, v := range runes2 {
		freq2 = append(freq2, v)
	}
	sort.Ints(freq2)

	for i, v := range freq1 {
		if v != freq2[i] {
			return false
		}
	}

	return true
}
