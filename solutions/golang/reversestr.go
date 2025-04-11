package golang

func reverseStr(s string, k int) string {
	runes := []rune(s)

	for group := 0; group < len(runes); group += Double * k {
		i := group
		j := Min(group+k, len(runes)) - 1

		for i < j {
			runes[i], runes[j] = runes[j], runes[i]
			i++
			j--
		}
	}

	return string(runes)
}
