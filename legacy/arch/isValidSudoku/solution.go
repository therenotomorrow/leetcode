package isValidSudoku

const sudokuSize = 9

type hs map[byte]struct{}

func isValidSudoku(board [][]byte) bool {
	rowsSet := make([]hs, sudokuSize)
	colsSet := make([]hs, sudokuSize)
	boxesSet := make([]hs, sudokuSize)

	for i := 0; i < sudokuSize; i++ {
		rowsSet[i] = make(hs)
		colsSet[i] = make(hs)
		boxesSet[i] = make(hs)
	}

	for i := 0; i < sudokuSize; i++ {
		for j := 0; j < sudokuSize; j++ {
			val := board[i][j]

			if val == '.' {
				continue
			}

			if _, ok := rowsSet[i][val]; ok {
				return false
			}
			rowsSet[i][val] = struct{}{}

			if _, ok := colsSet[j][val]; ok {
				return false
			}
			colsSet[j][val] = struct{}{}

			k := (i/3)*3 + j/3
			if _, ok := boxesSet[k][val]; ok {
				return false
			}
			boxesSet[k][val] = struct{}{}
		}
	}

	return true
}
