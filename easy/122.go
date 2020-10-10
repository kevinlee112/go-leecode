package easy




func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	buy, sell := -prices[0], 0

	for i := 1; i < len(prices); i++ {
		buy = max(buy, sell-prices[i])
		sell = max(buy+prices[i], sell)
	}
	return sell
}
