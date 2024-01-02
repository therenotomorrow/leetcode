package maxProfit

func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}

func maxProfit2(prices []int) int {
	profit := 0
	mPrice := 1_000_000

	for _, price := range prices {
		if price < mPrice {
			mPrice = price
		} else if price-mPrice > profit {
			profit = price - mPrice
		}
	}

	return profit
}
