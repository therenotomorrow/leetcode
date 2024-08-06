package golang

func minHeightShelves(books [][]int, shelfWidth int) int {
	var (
		memo    = make([][]int, len(books)+1) // alternative of `cache.NewCache()`
		dynamic func(curr int, remainW int, maxH int) int
	)

	for i := range memo {
		memo[i] = make([]int, shelfWidth+1)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	dynamic = func(curr int, remainW int, maxH int) int {
		// order of books is important that's why we just try to fit by width without optimize by height
		if val := memo[curr][remainW]; val != -1 {
			return val
		}

		currW := books[curr][0]
		currH := books[curr][1]

		newMaxH := Max(maxH, currH)

		if curr == len(books)-1 {
			if remainW >= currW {
				return newMaxH
			}

			return maxH + currH
		}

		// place on new shelf
		maxH += dynamic(curr+1, shelfWidth-currW, currH)

		if remainW >= currW {
			// place on current shelf
			maxH = Min(maxH, dynamic(curr+1, remainW-currW, newMaxH))
		}

		memo[curr][remainW] = maxH

		return maxH
	}

	return dynamic(0, shelfWidth, 0)
}
