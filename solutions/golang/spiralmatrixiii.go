package golang

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	const two = 2

	var (
		i                     = rStart
		j                     = cStart
		k                     = 0
		direction             = 0
		times                 = 1
		timesDirectionChanged = 0
		dirs                  = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		res                   = make([][]int, rows*cols)
	)

	for steps := 1; k < rows*cols; steps++ {
		if 0 <= i && i < rows && 0 <= j && j < cols {
			res[k] = []int{i, j}
			k++
		}

		i += dirs[direction][0]
		j += dirs[direction][1]

		if steps >= times {
			direction++
			direction %= 4
			steps = 0
			timesDirectionChanged++
		}

		if timesDirectionChanged >= two {
			timesDirectionChanged = 0
			times++
		}
	}

	return res
}
