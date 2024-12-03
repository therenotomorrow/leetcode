package golang

import (
	"strings"
)

func addSpaces(text string, spaces []int) string {
	newS := strings.Builder{}

	for i, j := 0, 0; i < len(text); i++ {
		if j < len(spaces) && i == spaces[j] {
			newS.WriteByte(' ')

			j++
		}

		newS.WriteByte(text[i])
	}

	return newS.String()
}
