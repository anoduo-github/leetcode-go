package question

//1748. 唯一元素的和
func sumOfUnique(nums []int) int {
	m := make(map[int]int, 0)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	res := 0
	for _, v := range nums {
		if m[v] == 1 {
			res = res + v
		}
	}
	return res
}
