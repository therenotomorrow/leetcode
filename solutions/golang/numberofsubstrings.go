package golang

func numberOfSubstrings(text string) int {
	var (
		aLet, bLet, cLet, ans, left int
		inc                         = func(r rune, val int) {
			switch r {
			case 'a':
				aLet += val
			case 'b':
				bLet += val
			case 'c':
				cLet += val
			}
		}
	)

	for right, r := range text {
		inc(r, 1)

		for aLet > 0 && bLet > 0 && cLet > 0 {
			ans += len(text) - right

			inc(rune(text[left]), -1)

			left++
		}
	}

	return ans
}
