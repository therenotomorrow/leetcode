package golang

func findTheLongestSubstring(text string) int {
	const (
		one = 1 << iota
		two
		four
		eight
		sixteen
	)

	var (
		state  = 0
		maxLen = 0
		masks  = map[rune]int{'a': one, 'e': two, 'i': four, 'o': eight, 'u': sixteen}
		states = map[int]int{0: -1}
	)

	for end, runa := range text {
		if mask, ok := masks[runa]; ok {
			state ^= mask
		}

		start, ok := states[state]
		if ok {
			maxLen = Max(maxLen, end-start)
		} else {
			states[state] = end
		}
	}

	return maxLen
}
