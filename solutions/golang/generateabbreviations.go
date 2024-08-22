package golang

import "strconv"

func generateAbbreviations(word string) []string {
	var (
		ans     = make([]string, 0)
		counter = func(word []rune, cnt int) []rune {
			if cnt > 0 {
				return []rune(string(word) + strconv.Itoa(cnt))
			}

			return word
		}
		backtrack func(curr []rune, i int, cnt int)
	)

	backtrack = func(curr []rune, i int, cnt int) {
		if i >= len(word) {
			ans = append(ans, string(counter(curr, cnt)))

			return
		}

		currLen := len(curr)

		curr = counter(curr, cnt)
		curr = append(curr, rune(word[i]))

		backtrack(curr, i+1, 0)

		curr = curr[:currLen]

		backtrack(curr, i+1, cnt+1)
	}

	backtrack(nil, 0, 0)

	return ans
}
