package minimumCardPickup

import "math"

func minimumCardPickup(cards []int) int {
	cardsMap := make(map[int]int)
	min := math.MaxInt

	for i, card := range cards {
		last, ok := cardsMap[card]

		if ok {
			min = mini(min, i-last+1)
		}

		cardsMap[card] = i
	}

	if min == math.MaxInt {
		min = -1
	}

	return min
}

func mini(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
