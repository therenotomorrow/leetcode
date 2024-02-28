package golang

func lengthOfLongestSubstringTwoDistinct(s string) int {
	var (
		maxLen, count int
		alphabet      = make([]byte, 255)
	)

	for l, r := 0, 0; r < len(s); r++ {
		rChar := s[r] - 'a'

		if alphabet[rChar] == 0 {
			count++
		}

		alphabet[rChar]++

		for count > 2 {
			lChar := s[l] - 'a'

			alphabet[lChar]--
			if alphabet[lChar] == 0 {
				count--
			}

			l++
		}

		maxLen = Max(maxLen, r-l+1)
	}

	return maxLen
}
