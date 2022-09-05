package question

import s1 "sort"

//322. 零钱兑换
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	//排序
	s1.Ints(coins)
	//记录结果
	res := 0
	//从右往左遍历
	for i := len(coins) - 1; i >= 0; i-- {
		x := amount / coins[i] //取余
		y := amount % coins[i] //取模
		if x == 0 {
			//表示amount小于最大的硬币值
			continue
		} else {
			if y > 0 {
				//表示此时不够除，还剩点
				res += x
				amount = y
				continue
			} else {
				//表示此时刚好除尽
				res += x
				amount = 0
				break
			}
		}
	}
	if amount != 0 {
		return -1
	}
	return res
}
