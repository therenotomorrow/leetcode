package golang

func numberOfPatterns(minPass int, maxPass int) int { //nolint:funlen,cyclop
	const (
		startLen = 1
		gridSize = 3
	)

	type gridT [gridSize][gridSize]bool

	var (
		directMoves = [][]int{
			{0, 1},
			{0, -1},
			{1, 0},
			{-1, 0},
			{1, 1},
			{-1, 1},
			{1, -1},
			{-1, -1},
			{-2, 1},
			{-2, -1},
			{2, 1},
			{2, -1},
			{1, -2},
			{-1, -2},
			{1, 2},
			{-1, 2},
		}
		skipMidMoves = [][]int{{0, 2}, {0, -2}, {2, 0}, {-2, 0}, {-2, -2}, {2, 2}, {2, -2}, {-2, 2}}
		backtrack    func(i int, j int, passLen int, grid gridT) int
	)

	backtrack = func(i int, j int, passLen int, grid gridT) int {
		if passLen > maxPass {
			return 0
		}

		if 0 > i || i >= gridSize || 0 > j || j >= gridSize || grid[i][j] {
			return 0
		}

		patterns := 0
		if passLen >= minPass {
			patterns = 1
		}

		grid[i][j] = true

		for _, move := range directMoves {
			patterns += backtrack(i+move[0], j+move[1], passLen+1, grid)
		}

		for _, move := range skipMidMoves {
			newRow := i + move[0]
			newCol := j + move[1]

			if 0 <= newRow && newRow < gridSize && 0 <= newCol && newCol < gridSize && !grid[newRow][newCol] {
				// check middle dot
				if !grid[i+move[0]/2][j+move[1]/2] {
					continue
				}

				patterns += backtrack(newRow, newCol, passLen+1, grid)
			}
		}

		grid[i][j] = false

		return patterns
	}

	// 4 corners + 4 middle of borders + 1 middle
	return backtrack(0, 0, startLen, gridT{})*4 + backtrack(0, 1, startLen, gridT{})*4 + backtrack(1, 1, startLen, gridT{})
}
