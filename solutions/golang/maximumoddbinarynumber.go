package golang

import "strings"

func maximumOddBinaryNumber(str string) string {
	cnt := 0

	for _, r := range str {
		if r == '1' {
			cnt++
		}
	}

	var builder strings.Builder

	// look at the structure of max odd numbers
	builder.WriteString(strings.Repeat("1", cnt-1))
	builder.WriteString(strings.Repeat("0", len(str)-cnt))
	builder.WriteString("1")

	return builder.String()
}
