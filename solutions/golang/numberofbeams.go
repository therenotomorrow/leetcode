package golang

func numberOfBeams(bank []string) int {
	ans, prevCnt := 0, 0

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
