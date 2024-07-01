package golang

func removeDigit(number string, digit byte) string {
	ans := ""

	for i := range len(number) {
		if number[i] != digit {
			continue
		}

		tmp := number[:i] + number[i+1:]
		if tmp > ans {
			ans = tmp
		}
	}

	return ans
}
