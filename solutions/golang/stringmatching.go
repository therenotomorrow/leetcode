package golang

import "strings"

func stringMatching(words []string) []string {
	set := NewSet[string]()

	for _, word := range words {
		for _, search := range words {
			if word == search {
				continue
			}

			if strings.Contains(word, search) {
				set.Add(search)
			}
		}
	}

	return set.Values()
}
