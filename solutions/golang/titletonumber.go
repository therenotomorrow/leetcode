package golang

func titleToNumber(columnTitle string) int {
	num := 0

	for _, r := range columnTitle {
		num *= 26
		num += int(r-'A') + 1
	}

	return num
}
