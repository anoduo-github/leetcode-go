package offer

//剑指 Offer 63. 股票的最大利润
func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			temp := prices[j] - prices[i]
			if temp > max {
				max = temp
			}
		}
	}
	return max
}
