package golang

func maxEqualRowsAfterFlips(matrix [][]int) int {
	patterns := make(map[string]int)

	for _, row := range matrix {
		pattern := make([]rune, len(row))

		for i, cell := range row {
			if row[0] == cell {
				pattern[i] = 'T'
			} else {
				pattern[i] = 'F'
			}
		}

		patterns[string(pattern)]++
	}

	cnt := 0

	for _, times := range patterns {
		cnt = Max(cnt, times)
	}

	return cnt
}
