package main

func main() {

}

// 121 买股票的最佳时机
func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else if p > maxProfit {
			maxProfit = p - minPrice
		}
	}

	return maxProfit
}
