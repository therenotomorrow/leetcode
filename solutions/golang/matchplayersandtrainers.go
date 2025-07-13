package golang

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)

	ans := 0

	for _, trainer := range trainers {
		if trainer >= players[ans] {
			ans++
		}

		if ans >= len(players) {
			break
		}
	}

	return ans
}
