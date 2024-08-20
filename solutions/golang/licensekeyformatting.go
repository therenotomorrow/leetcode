package golang

import (
	"strings"
	"unicode"
)

func licenseKeyFormatting(s string, k int) string {
	runes := make([]rune, 0)

	for _, runa := range s {
		if runa == '-' {
			continue
		}

		runes = append(runes, unicode.ToUpper(runa))
	}

	key := strings.Builder{}

	for i, runa := range runes {
		if (i-len(runes)%k)%k == 0 && key.Len() > 0 {
			key.WriteRune('-')
		}

		key.WriteRune(runa)
	}

	return key.String()
}
