package question

//1475. 商品折扣后的最终价格
func finalPrices(prices []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(prices); i++ {
		tag := true
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				res = append(res, prices[i]-prices[j])
				tag = false
				break
			}
		}
		if tag {
			res = append(res, prices[i])
		}
	}
	return res
}
