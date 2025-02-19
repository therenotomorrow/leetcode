package golang

import "sort"

func getHappyString(n int, k int) string {
	var (
		strings   []string
		backtrack func(curr string)
	)

	backtrack = func(curr string) {
		if len(curr) == n {
			strings = append(strings, curr)

			return
		}

		for _, letter := range []byte{'a', 'b', 'c'} {
			if l := len(curr); l > 0 && curr[l-1] == letter {
				continue
			}

			backtrack(curr + string(letter))
		}
	}

	backtrack("")

	if len(strings) < k {
		return ""
	}

	sort.Strings(strings)

	return strings[k-1]
}
