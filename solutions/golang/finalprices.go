package golang

func finalPrices(prices []int) []int {
	for i, price := range prices {
		for _, next := range prices[i+1:] {
			if price >= next {
				prices[i] = price - next

				break
			}
		}
	}

	return prices
}
