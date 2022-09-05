package offer

//剑指 Offer 03. 数组中重复的数字
func findRepeatNumber(nums []int) int {
	m := make(map[int]int, 0)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		} else {
			m[v] = 1
		}
	}
	return -1
}
