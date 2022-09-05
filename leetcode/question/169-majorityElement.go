package question

//169. 多数元素
func majorityElement(nums []int) int {
	target := len(nums) / 2
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for k, i := range m {
		if i > target {
			return k
		}
	}
	return -1
}
