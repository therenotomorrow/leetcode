package golang

import "sort"

func maxCoins(piles []int) int {
	const (
		turn = 2
		size = 3
	)

	sort.Ints(piles)

	// we took the coin after the last maximum
	hand := len(piles) - turn
	selfCoins := 0

	for turns := len(piles) / size; turns > 0; turns-- {
		selfCoins += piles[hand]
		hand -= turn
	}

	return selfCoins
}
