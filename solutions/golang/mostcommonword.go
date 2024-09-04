package golang

import (
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	re := regexp.MustCompile(`\w+`)
	cnt := make(map[string]int)

	for _, word := range re.FindAllString(paragraph, -1) {
		cnt[strings.ToLower(word)]++
	}

	for _, word := range banned {
		delete(cnt, word)
	}

	mostCommon := ""
	mostTimes := 0

	for word, times := range cnt {
		if times > mostTimes {
			mostTimes = times
			mostCommon = word
		}
	}

	return mostCommon
}
