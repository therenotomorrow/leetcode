package golang

func commonChars(words []string) []string {
	origin := make([]rune, Alphabet)
	for _, r := range words[0] {
		origin[r-'a']++
	}

	for _, word := range words[1:] {
		tmp := make([]rune, Alphabet)

		for _, r := range word {
			tmp[r-'a']++
		}

		for i := range Alphabet {
			if origin[i] >= tmp[i] {
				origin[i] = tmp[i]
			}
		}
	}

	ans := make([]string, 0)

	for i, times := range origin {
		for ; times > 0; times-- {
			ans = append(ans, string(rune(i)+'a'))
		}
	}

	return ans
}
