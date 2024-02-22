package golang

import "sort"

func maxCoins(piles []int) int {
	sort.Ints(piles)

	// we took the coin after the last maximum
	hand := len(piles) - 2
	selfCoins := 0

	for turns := len(piles) / 3; turns > 0; turns-- {
		selfCoins += piles[hand]
		hand -= 2
	}

	return selfCoins
}
