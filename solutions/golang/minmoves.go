package golang

func minMoves(rooks [][]int) int {
	countSort := [][]int{make([]int, len(rooks)), make([]int, len(rooks))}
	for _, rook := range rooks {
		countSort[0][rook[0]]++
		countSort[1][rook[1]]++
	}

	moves := 0
	rowsDiff := 0
	colsDiff := 0

	for i := range rooks {
		rowsDiff += countSort[0][i] - 1
		colsDiff += countSort[1][i] - 1

		moves += Abs(rowsDiff) + Abs(colsDiff)
	}

	return moves
}
