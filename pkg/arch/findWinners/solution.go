package findWinners

import "sort"

func findWinners(matches [][]int) [][]int {
	games := make(map[int]int)

	for _, match := range matches {
		win, los := match[0], match[1]

		games[win] += 0
		games[los] -= 1
	}

	wins := make([]int, 0)
	loss := make([]int, 0)

	for player, game := range games {
		if game == 0 {
			wins = append(wins, player)
		}
		if game == -1 {
			loss = append(loss, player)
		}
	}

	sort.Ints(wins)
	sort.Ints(loss)

	return [][]int{wins, loss}
}
