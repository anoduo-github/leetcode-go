package question

//1423. 可获得的最大点数
func maxScore(cardPoints []int, k int) int {
	//获取数组长度
	l := len(cardPoints)
	//求数组总大小
	total := 0
	for _, v := range cardPoints {
		total += v
	}
	//滑动窗口，从左起求第一次窗口数之和sum
	sum := 0
	for i := 0; i < l-k; i++ {
		sum += cardPoints[i]
	}
	//第一次的最大值max
	max := total - sum
	//窗口朝滑动
	for i := l - k; i < l; i++ {
		sum = sum + cardPoints[i] - cardPoints[i-(l-k)]
		max = getMax(max, total-sum)
	}
	return max
}
