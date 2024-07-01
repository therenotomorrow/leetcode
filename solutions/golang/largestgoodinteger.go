package golang

import "strings"

func largestGoodInteger(num string) string {
	const size = 3

	var (
		prev, curr rune
		cnt        int
	)

	for _, r := range num {
		if r != curr {
			curr = r
			cnt = 1
		} else if r == curr {
			cnt++
		}

		if cnt == size {
			if curr > prev {
				prev = curr
			}
		}
	}

	if prev == rune(0) {
		return ""
	}

	return strings.Repeat(string(prev), size)
}
