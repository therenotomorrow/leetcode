package solutions

import "sort"

func closeStrings(word1 string, word2 string) bool {
	// step 0: check simple condition
	if len(word1) != len(word2) {
		return false
	}

	// step 1: check that all characters the same
	runes1 := make([]int, 26)
	for _, r := range word1 {
		runes1[r-'a']++
	}

	runes2 := make([]int, 26)
	for _, r := range word2 {
		runes2[r-'a']++

		if runes1[r-'a'] == 0 {
			return false
		}
	}

	// step 2: check that all frequencies the same
	sort.Ints(runes1)
	sort.Ints(runes2)

	for i := 0; i < 26; i++ {
		if runes1[i] != runes2[i] {
			return false
		}
	}

	return true
}
