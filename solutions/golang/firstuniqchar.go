package golang

func firstUniqChar(str string) int {
	set := make(map[rune]int)
	for _, r := range str {
		set[r]++
	}

	for i, r := range str {
		if set[r] == 1 {
			return i
		}
	}

	return -1
}
