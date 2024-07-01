package golang

func lengthOfLongestSubstringTwoDistinct(str string) int {
	var (
		maxLen, count int
		alphabet      = make([]byte, Byte)
	)

	for left, right := 0, 0; right < len(str); right++ {
		rChar := str[right] - 'a'

		if alphabet[rChar] == 0 {
			count++
		}

		alphabet[rChar]++

		for count > 2 {
			lChar := str[left] - 'a'

			alphabet[lChar]--
			if alphabet[lChar] == 0 {
				count--
			}

			left++
		}

		maxLen = Max(maxLen, right-left+1)
	}

	return maxLen
}
