package golang

func stringShift(text string, shift [][]int) string {
	var leftShift, rightShift int

	for _, pair := range shift {
		if pair[0] == 0 {
			leftShift += pair[1]
		} else {
			rightShift += pair[1]
		}
	}

	diff := rightShift - leftShift
	if diff <= 0 {
		diff = -diff
		diff %= len(text)
	} else {
		diff %= len(text)
		diff = len(text) - diff
	}

	return text[diff:] + text[:diff]
}
