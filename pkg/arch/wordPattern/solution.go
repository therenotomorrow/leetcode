package wordPattern

import "strings"

func wordPattern(pattern string, s string) bool {
	parts := strings.Fields(s)

	if len(pattern) != len(parts) {
		return false
	}

	map1 := make(map[rune]string)
	map2 := make(map[string]rune)

	for i, r := range pattern {
		_, ok1 := map1[r]
		_, ok2 := map2[parts[i]]

		if !ok1 {
			map1[r] = parts[i]
		}
		if !ok2 {
			map2[parts[i]] = r
		}

		if map1[r] != parts[i] || map2[parts[i]] != r {
			return false
		}
	}

	return true
}
