package solutions

import "sort"

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)

	if diff := prices[0] + prices[1]; diff <= money {
		money -= diff
	}

	return money
}
