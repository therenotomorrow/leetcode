package golang

import "strings"

func defangIPaddr(address string) string {
	var builder strings.Builder

	for _, runa := range address {
		if runa == '.' {
			builder.WriteString("[.]")
		} else {
			builder.WriteRune(runa)
		}
	}

	return builder.String()
}
