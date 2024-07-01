package golang

func isAnagram(str string, tpl string) bool {
	var usedS [26]rune
	for _, r := range str {
		usedS['z'-r]++
	}

	var usedT [26]rune
	for _, r := range tpl {
		usedT['z'-r]++
	}

	return usedS == usedT
}
