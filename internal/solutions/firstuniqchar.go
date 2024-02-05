package solutions

func firstUniqChar(s string) int {
	set := make(map[rune]int)
	for _, r := range s {
		set[r]++
	}

	for i, r := range s {
		if set[r] == 1 {
			return i
		}
	}

	return -1
}
