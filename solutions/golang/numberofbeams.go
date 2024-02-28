package golang

func numberOfBeams(bank []string) int {
	var ans, prevCnt int

	for _, row := range bank {
		currCnt := 0

		for _, r := range row {
			if r == '1' {
				currCnt++
			}
		}

		if currCnt > 0 {
			ans += prevCnt * currCnt
			prevCnt = currCnt
		}
	}

	return ans
}
