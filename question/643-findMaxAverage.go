package question

//643.子数组最大平均数 I
func findMaxAverage(nums []int, k int) float64 {
	//获取数组总长
	l := len(nums)
	//滑动窗口求解，长度为k,求第一次从左起的窗口平均值
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ave := float64(sum) / float64(k)
	for i := k; i < l; i++ {
		sum = sum + nums[i] - nums[i-k]
		ave = getMaxave(ave, float64(sum)/float64(k))
	}
	return ave
}

func getMaxave(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}
