package golang

func balancedStringSplit(s string) int {
	cnt := 0
	balanced := 0

	for _, runa := range s {
		if runa == 'L' {
			balanced--
		} else {
			balanced++
		}

		if balanced == 0 {
			cnt++
		}
	}

	return cnt
}
