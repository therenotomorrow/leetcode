package golang

import "strings"

func countSegments(s string) int {
	return len(strings.Fields(s))
}
