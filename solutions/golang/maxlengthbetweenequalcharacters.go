package golang

func maxLengthBetweenEqualCharacters(str string) int {
	indexes := make(map[rune]int)
	maxLen := -1

	for idx, runa := range str {
		j, ok := indexes[runa]

		if ok {
			maxLen = Max(maxLen, idx-j-1)
		} else {
			indexes[runa] = idx
		}
	}

	return maxLen
}
