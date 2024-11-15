package golang

func checkOnesSegment(s string) bool {
	last := '1'

	for _, curr := range s {
		if curr == '1' && last == '0' {
			return false
		}

		last = curr
	}

	return true
}
