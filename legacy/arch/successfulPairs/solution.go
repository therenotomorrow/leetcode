package successfulPairs

import (
	"sort"
)

func binarySearch(potions []int, spell int, success int64) int {
	l, r, m := 0, len(potions)-1, 0

	for l <= r {
		m = (l + r) / 2

		if int64(spell)*int64(potions[m]) < success {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return l
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)

	n, m := len(potions), len(spells)
	ans := make([]int, m)

	for i, spell := range spells {
		ans[i] = n - binarySearch(potions, spell, success)
	}

	return ans
}
