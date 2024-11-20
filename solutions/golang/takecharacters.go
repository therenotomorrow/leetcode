package golang

func takeCharacters(text string, k int) int {
	const size = 3

	cnt := make([]int, size)

	for _, r := range text {
		cnt[r-'a']++
	}

	for _, t := range cnt {
		if t < k {
			return -1
		}
	}

	i := 0
	curr := make([]int, size)
	maxLen := 0

	for j, r := range text {
		curr[r-'a']++

		for i <= j && (cnt[0]-curr[0] < k || cnt[1]-curr[1] < k || cnt[2]-curr[2] < k) {
			curr[text[i]-'a']--

			i++
		}

		maxLen = Max(maxLen, j-i+1)
	}

	return len(text) - maxLen
}
