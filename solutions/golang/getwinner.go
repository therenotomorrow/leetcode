package golang

func getWinner(arr []int, k int) int {
	winner := arr[0]
	streak := 0

	// at the end of arr we will find max
	for _, opponent := range arr[1:] {
		if streak == k {
			break
		}

		if opponent < winner {
			streak++

			continue
		}

		winner = opponent
		streak = 1
	}

	return winner
}
