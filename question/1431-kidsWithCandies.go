package question

//1431. 拥有最多糖果的孩子
func kidsWithCandies(candies []int, extraCandies int) []bool {
	//1.先找出最大的
	max := -1
	for _, v := range candies {
		if v >= max {
			max = v
		}
	}
	//2.遍历数据加3
	res := make([]bool, 0)
	for _, v := range candies {
		if v+extraCandies >= max {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}
