package golang

import "slices"

func convertToTitle(columnNumber int) string {
	title := make([]byte, 0)

	for columnNumber > 0 {
		columnNumber-- // to make number as true base 26

		title = append(title, byte(columnNumber%Alphabet)+'A')

		columnNumber /= Alphabet
	}

	slices.Reverse(title)

	return string(title)
}
