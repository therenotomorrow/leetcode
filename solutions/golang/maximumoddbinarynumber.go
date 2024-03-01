package golang

import "strings"

func maximumOddBinaryNumber(s string) string {
	cnt := 0

	for _, r := range s {
		if r == '1' {
			cnt++
		}
	}

	var sb strings.Builder

	// look at the structure of max odd numbers
	sb.WriteString(strings.Repeat("1", cnt-1))
	sb.WriteString(strings.Repeat("0", len(s)-cnt))
	sb.WriteString("1")

	return sb.String()
}
