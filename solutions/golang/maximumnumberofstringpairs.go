package golang

func maximumNumberOfStringPairs(words []string) int {
	ans := 0
	cnt := make(map[string]struct{})

	for _, word := range words {
		// took from reverseonlyletters.go: https://leetcode.com/problems/reverse-only-letters/description/
		reversed := reverseOnlyLetters(word)
		if word == reversed {
			continue
		}

		cnt[word] = struct{}{}
		if _, ok := cnt[reversed]; ok {
			ans++
		}
	}

	return ans
}
