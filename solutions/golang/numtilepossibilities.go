package golang

func numTilePossibilities(tiles string) int {
	var (
		counter func() int
		runes   = make([]rune, Alphabet)
	)

	for _, runa := range tiles {
		runes[runa-'A']++
	}

	counter = func() int {
		cnt := 0

		for idx := range runes {
			if runes[idx] == 0 {
				continue
			}

			cnt++
			runes[idx]--
			cnt += counter()
			runes[idx]++
		}

		return cnt
	}

	return counter()
}
