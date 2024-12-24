package golang

func isBalanced2(num string) bool {
	var evens, odds int

	for i, r := range num {
		number := int(r - '0')

		if i%2 == 0 {
			evens += number
		} else {
			odds += number
		}
	}

	return evens == odds
}
