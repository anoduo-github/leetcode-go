package question

//217. 存在重复元素
func containsDuplicate(nums []int) bool {
	m := make(map[int]int, 0)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = 1
		}
	}
	return false
}
