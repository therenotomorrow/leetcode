package golang

import "strings"

func splitWordsBySeparator(words []string, separator byte) []string {
	ans := make([]string, 0)

	for _, word := range words {
		for _, sword := range strings.Split(word, string(separator)) {
			if len(sword) > 0 {
				ans = append(ans, sword)
			}
		}
	}

	return ans
}
