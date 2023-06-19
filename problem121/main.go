package problem121

func maxProfit(prices []int) int {
	min := prices[0]
	var profit int
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}
	return profit
}
