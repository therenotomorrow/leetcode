package longestCommonPrefix

import "strings"

func longestCommonPrefix(strs []string) string {
	var prefix strings.Builder

	min := 200
	for _, str := range strs {
		if len(str) < min {
			min = len(str)
		}
	}

	for i := 0; i < min; i++ {
		pattern := strs[0][i]

		for _, str := range strs {
			if pattern != str[i] {
				return prefix.String()
			}
		}

		prefix.WriteByte(pattern)
	}

	return prefix.String()
}
