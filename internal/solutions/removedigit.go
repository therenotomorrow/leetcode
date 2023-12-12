package solutions

func removeDigit(number string, digit byte) string {
	ans := ""

	for i := 0; i < len(number); i++ {
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
