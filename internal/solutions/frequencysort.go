package solutions

import (
	"strings"

	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func frequencySort(s string) string {
	maxCnt := 0
	cnt := make(map[rune]int)

	for _, r := range s {
		cnt[r]++

		maxCnt = mathfunc.Max(maxCnt, cnt[r])
	}

	buckets := make([][]rune, maxCnt+1)
	for r, times := range cnt {
		buckets[times] = append(buckets[times], r)
	}

	sb := strings.Builder{}

	for i := maxCnt; i >= 0; i-- {
		for _, r := range buckets[i] {
			for j := 0; j < i; j++ {
				sb.WriteRune(r)
			}
		}
	}

	return sb.String()
}
