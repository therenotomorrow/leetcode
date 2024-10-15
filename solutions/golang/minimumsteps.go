package golang

func minimumSteps(s string) int64 {
	cnt := 0
	ones := 0

	for _, r := range s {
		if r == '1' {
			ones++
		} else {
			// because we swap like 1110 -> 1101 -> 1011 -> 0111 == 3
			cnt += ones
		}
	}

	return int64(cnt)
}
