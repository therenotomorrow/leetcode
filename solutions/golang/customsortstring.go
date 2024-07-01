package golang

import "strings"

func customSortString(order string, str string) string {
	cnt := make(map[rune]int)
	for _, runa := range str {
		cnt[runa]++
	}

	builder := strings.Builder{}

	for _, runa := range order {
		times := cnt[runa]

		for range times {
			builder.WriteRune(runa)
		}

		delete(cnt, runa)
	}

	for runa, times := range cnt {
		for range times {
			builder.WriteRune(runa)
		}
	}

	return builder.String()
}
