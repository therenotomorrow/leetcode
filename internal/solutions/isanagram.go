package solutions

func isAnagram(s string, t string) bool {
	var usedS [26]rune
	for _, r := range s {
		usedS['z'-r]++
	}

	var usedT [26]rune
	for _, r := range t {
		usedT['z'-r]++
	}

	return usedS == usedT
}
