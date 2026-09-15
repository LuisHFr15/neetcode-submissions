func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxProfit := 0
	buy := 0
	sell := 1

	for sell < len(prices) {
		for buy < sell && prices[buy] > prices[sell] {
			buy++
		}
		maxProfit = max(maxProfit, prices[sell] - prices[buy])
		sell++
	}
	return maxProfit
}
