package golang

func spiralMatrix(rows int, cols int, head *ListNode) [][]int { //nolint:cyclop
	var (
		i         = 0
		j         = 0
		k         = 0
		direction = 0
		dirs      = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		matrix    = make([][]int, rows)
	)

	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	var (
		left   = 0
		right  = cols
		upper  = 0
		bottom = rows
	)

	for curr := head; k < rows*cols; {
		if upper <= i && i < bottom && left <= j && j < right {
			if curr != nil {
				matrix[i][j] = curr.Val
				curr = curr.Next
			} else {
				matrix[i][j] = -1
			}

			k++
		}

		i += dirs[direction][0]
		j += dirs[direction][1]

		switch {
		case j >= right:
			j = right - 1
			upper++
			direction++
		case i >= bottom:
			i = bottom - 1
			right--
			direction++
		case j < left:
			j = left
			bottom--
			direction++
		case i < upper:
			i = upper
			left++
			direction++
		}

		direction %= 4
	}

	return matrix
}
