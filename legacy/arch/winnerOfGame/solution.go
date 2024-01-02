package winnerOfGame

func winnerOfGame(colors string) bool {
	var alice, bob int

	for i := 1; i < len(colors)-1; i++ {
		if colors[i-1] == colors[i] && colors[i] == colors[i+1] {
			if colors[i] == 'A' {
				alice++
			} else {
				bob++
			}
		}
	}

	return alice-bob > 0
}
