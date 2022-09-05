package primary

import "testing"

func TestMaxProfit(t *testing.T) {
	prices := []int{1, 2, 3, 1}
	res := rob(prices)
	t.Errorf("res: %d", res)
}
