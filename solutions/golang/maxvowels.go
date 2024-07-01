package golang

func maxVowels(str string, window int) int {
	vowels := map[byte]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
	maxCnt := 0

	for i := range str[:window] {
		maxCnt += vowels[str[i]]
	}

	curCnt := maxCnt

	for i, r := range str[window:] {
		curCnt -= vowels[str[i]]
		curCnt += vowels[byte(r)]
		maxCnt = Max(maxCnt, curCnt)
	}

	return maxCnt
}
