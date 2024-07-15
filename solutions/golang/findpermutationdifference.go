package golang

func findPermutationDifference(text string, tmpl string) int {
	cnt := make(map[rune]int)
	for i, runa := range text {
		cnt[runa] = i
	}

	diff := 0
	for i, runa := range tmpl {
		diff += Abs(i - cnt[runa])
	}

	return diff
}
