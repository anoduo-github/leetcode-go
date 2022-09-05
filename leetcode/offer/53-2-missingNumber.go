package offer

//剑指 Offer 53 - II. 0～n-1中缺失的数字
func missingNumber(nums []int) int {
	temp := make([]int, len(nums)+1)
	for _, v := range nums {
		temp[v] = 1
	}
	for k, v := range temp {
		if v == 0 {
			return k
		}
	}
	return -1
}
