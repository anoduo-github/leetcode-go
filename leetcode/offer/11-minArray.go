package offer

//剑指 Offer 11. 旋转数组的最小数字
func minArray(numbers []int) int {
	min := numbers[0]
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] && min > numbers[i+1] {
			min = numbers[i+1]
		}
	}
	return min
}
