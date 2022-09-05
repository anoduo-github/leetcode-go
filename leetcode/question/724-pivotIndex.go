package question

//724. 寻找数组的中心下标
func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	sum := 0
	for i, v := range nums {
		if sum*2+v == total {
			return i
		}
		sum += v
	}
	return -1
}
