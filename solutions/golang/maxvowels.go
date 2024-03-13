package golang

func maxVowels(s string, k int) int {
	vowels := map[byte]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
	maxCnt := 0

	for i := range s[:k] {
		maxCnt += vowels[s[i]]
	}

	curCnt := maxCnt

	for i, r := range s[k:] {
		curCnt -= vowels[s[i]]
		curCnt += vowels[byte(r)]
		maxCnt = Max(maxCnt, curCnt)
	}

	return maxCnt
}
