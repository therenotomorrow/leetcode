package nextGreatestLetter

func nextGreatestLetter(letters []byte, target byte) byte {
	l := 0
	r := len(letters) - 1

	for m := (l + r) / 2; l <= r; m = (l + r) / 2 {
		if letters[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if l == len(letters) {
		return letters[0]
	}

	return letters[l]
}
