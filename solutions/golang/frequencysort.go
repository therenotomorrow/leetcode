package golang

import (
	"strings"
)

func frequencySort(s string) string {
	maxCnt := 0
	cnt := make(map[rune]int)

	for _, r := range s {
		cnt[r]++

		maxCnt = Max(maxCnt, cnt[r])
	}

	buckets := make([][]rune, maxCnt+1)
	for r, times := range cnt {
		buckets[times] = append(buckets[times], r)
	}

	builder := strings.Builder{}

	for i := maxCnt; i >= 0; i-- {
		for _, r := range buckets[i] {
			for range i {
				builder.WriteRune(r)
			}
		}
	}

	return builder.String()
}
