package question

//1460. 通过翻转子数组使两个数组相等
func canBeEqual(target []int, arr []int) bool {
	temp := make([]int, 1001)
	for _, v := range target {
		temp[v]++
	}
	for _, v := range arr {
		temp[v]--
	}
	for _, v := range temp {
		if v != 0 {
			return false
		}
	}
	return true
}
