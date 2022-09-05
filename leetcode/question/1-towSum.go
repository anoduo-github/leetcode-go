package question

//1.两数之和
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if i, ok := m[target-v]; ok {
			return []int{i, k}
		} else {
			m[v] = k
		}
	}
	return nil
}
