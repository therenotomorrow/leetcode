package golang

func groupStrings(strings []string) [][]string {
	shift := func(s string) string {
		runes := []rune(s)
		dist := runes[0] - 'a'

		for i := range runes {
			runes[i] = rune(int(runes[i]-dist)%Alphabet) + 'a'
		}

		return string(runes)
	}

	groups := make(map[string][]string)

	for _, str := range strings {
		key := shift(str)
		groups[key] = append(groups[key], str)
	}

	ans := make([][]string, 0, len(groups))
	for _, words := range groups {
		ans = append(ans, words)
	}

	return ans
}
