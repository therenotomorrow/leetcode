package golang

func minOperations1(s string) int {
	cnt := 0 // will count '1'

	for i, digit := range s {
		if i%2 == 0 {
			if digit == '0' {
				cnt++
			}
		} else {
			if digit == '1' {
				cnt++
			}
		}
	}

	// lees correct '1' or '0'
	return Min(cnt, len(s)-cnt)
}
