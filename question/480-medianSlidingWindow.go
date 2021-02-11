package question

//480. 滑动窗口中位数
/* 超时了 */
func medianSlidingWindow(nums []int, k int) []float64 {
	//左指针, 右指针
	left, right := 0, k-1
	//结果集
	res := make([]float64, 0)
	for right < len(nums) {
		arr := make([]int, 0)
		for i := left; i <= right; i++ {
			arr = append(arr, nums[i])
		}
		res = append(res, getAvg(arr))
		left++
		right++
	}
	return res
}

func getAvg(nums []int) float64 {
	//先排序
	sort(nums)
	//获取长度
	l := len(nums)
	if l%2 == 0 {
		//偶数
		return float64(nums[l/2-1]+nums[l/2]) / 2
	}
	return float64(nums[l/2])
}

func sort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
