package golang

import (
	"unicode"
)

func replaceDigits(s string) string {
	res := []rune(s)

	for i, r := range res {
		if unicode.IsDigit(r) {
			r = res[i-1] + (r - '0')
		}

		res[i] = r
	}

	return string(res)
}
