
func maxProfit(prices []int) int {
	lowest := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		price := prices[i]
		if price < lowest {
			lowest = price
		} else {
			profit := price - lowest

			if profit > maxProfit {
				maxProfit = profit
			}

		}
	}

	return maxProfit

}