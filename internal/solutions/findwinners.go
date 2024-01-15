package solutions

import "sort"

func findWinners(matches [][]int) [][]int {
	games := make(map[int]int)

	for _, match := range matches {
		games[match[0]] += 0
		games[match[1]]--
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
