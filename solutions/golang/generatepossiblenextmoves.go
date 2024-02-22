package golang

func generatePossibleNextMoves(currentState string) []string {
	moves := make([]string, 0)

	for i := 0; i < len(currentState)-1; i++ {
		if currentState[i] == '+' && currentState[i+1] == '+' {
			moves = append(moves, currentState[:i]+"--"+currentState[i+2:])
		}
	}

	return moves
}
