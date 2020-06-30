package easy

func FinalPrices(prices []int) []int {
	lens := len(prices)
	for i:=0;i<lens;i++ {
		for j:=i+1;j<lens;j++ {
			if prices[j] < prices[i] {
				prices[i] = prices[i] - prices[j]
				continue
			}
		}
	}

	return prices
}