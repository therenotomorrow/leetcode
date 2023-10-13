package firstUniqChar

func firstUniqChar(s string) int {
	set := make(map[rune]int)
	for _, r := range s {
		set[r]++
	}

	for i, r := range s {
		val := set[r]
		if val == 1 {
			return i
		}
	}

	return -1
}
