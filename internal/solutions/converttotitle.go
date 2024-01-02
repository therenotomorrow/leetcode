package solutions

import "slices"

func convertToTitle(columnNumber int) string {
	title := make([]byte, 0)

	for columnNumber > 0 {
		columnNumber-- // to make number as true base 26

		title = append(title, byte(columnNumber%26)+'A')

		columnNumber /= 26
	}

	slices.Reverse(title)

	return string(title)
}
