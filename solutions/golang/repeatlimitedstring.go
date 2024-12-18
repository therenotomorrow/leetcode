package golang

import (
	"strings"
)

func repeatLimitedString(s string, repeatLimit int) string {
	arr := make([]int, Alphabet)
	for _, runa := range s {
		arr[runa-'a']++
	}

	ans := strings.Builder{}

	for idx := Alphabet - 1; idx >= 0; {
		if arr[idx] == 0 {
			idx--

			continue
		}

		times := Min(arr[idx], repeatLimit)
		arr[idx] -= times

		for range times {
			ans.WriteRune(rune(idx + 'a'))
		}

		if arr[idx] <= 0 {
			continue
		}

		next := idx - 1

		for next >= 0 && arr[next] == 0 {
			next--
		}

		if next < 0 {
			break
		}

		ans.WriteRune(rune('a' + next))

		arr[next]--
	}

	return ans.String()
}
