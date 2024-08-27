package golang

import "strings"

func findWords(words []string) []string {
	rows := []Set[string]{NewSet[string](), NewSet[string](), NewSet[string]()}

	rows[0].Populate("q", "w", "e", "r", "t", "y", "u", "i", "o", "p")
	rows[1].Populate("a", "s", "d", "f", "g", "h", "j", "k", "l")
	rows[2].Populate("z", "x", "c", "v", "b", "n", "m")

	ans := make([]string, 0)

	for _, word := range words {
		for _, set := range rows {
			found := true

			for _, r := range strings.ToLower(word) {
				if !set.Contains(string(r)) {
					found = false

					break
				}
			}

			if found {
				ans = append(ans, word)

				break
			}
		}
	}

	return ans
}
