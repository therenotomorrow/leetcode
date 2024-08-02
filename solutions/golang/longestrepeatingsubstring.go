package golang

func longestRepeatingSubstring(text string) int {
	var set Set[string]

	for maxLen := len(text) - 1; maxLen > 0; maxLen-- {
		set = NewSet[string]()

		for i := range len(text) + 1 - maxLen {
			currS := text[i : i+maxLen]

			if set.Contains(currS) {
				return maxLen
			}

			set.Add(currS)
		}
	}

	return 0
}
