func maxProfit(prices []int) int {
	result := 0
	buy := prices[0]
	sell := buy
	for _, value := range prices{
		if value < buy {
			buy = value
			sell = value
		}
		if value > sell {
			sell = value
		}
		if sell - buy > result{
			result = sell - buy
		}
	}

	return result
}
