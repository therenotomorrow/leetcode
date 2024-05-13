package golang

func exist(board [][]byte, word string) bool {
	var (
		rowDirs   = []int{1, 0, -1, 0}
		colDirs   = []int{0, 1, 0, -1}
		backtrack func(row int, col int, idx int) bool
	)

	backtrack = func(row int, col int, idx int) bool {
		if idx >= len(word) {
			return true
		}

		if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || board[row][col] != word[idx] {
			return false
		}

		board[row][col] = '0'

		for i := range rowDirs {
			if backtrack(row+rowDirs[i], col+colDirs[i], idx+1) {
				return true
			}
		}

		board[row][col] = word[idx]

		return false
	}

	for i, rows := range board {
		for j := range rows {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}
