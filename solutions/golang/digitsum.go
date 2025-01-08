package golang

import (
	"slices"
	"strconv"
	"strings"
)

func digitSum(text string, parts int) string {
	for len(text) > parts {
		var update strings.Builder

		for group := range slices.Chunk([]rune(text), parts) {
			num := 0

			for _, digit := range group {
				num += int(digit - '0')
			}

			update.WriteString(strconv.Itoa(num))
		}

		text = update.String()
	}

	return text
}
