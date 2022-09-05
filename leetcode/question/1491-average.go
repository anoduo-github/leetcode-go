package question

//1491. 去掉最低工资和最高工资后的工资平均值
func average(salary []int) float64 {
	if len(salary) == 0 {
		return 0
	}
	max := salary[0]
	min := salary[0]
	res := 0
	for _, v := range salary {
		res += v
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	res = res - max - min
	return float64(res) / float64(len(salary)-2)
}
