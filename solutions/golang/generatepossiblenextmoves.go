package golang

func generatePossibleNextMoves(currentState string) []string {
	moves := make([]string, 0)

	for i := range len(currentState) - 1 {
		if currentState[i] == '+' && currentState[i+1] == '+' {
			moves = append(moves, currentState[:i]+"--"+currentState[i+2:])
		}
	}

	return moves
}
