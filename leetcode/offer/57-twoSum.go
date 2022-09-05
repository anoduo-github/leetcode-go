package offer

//剑指 Offer 57. 和为s的两个数字
func twoSum(nums []int, target int) []int {
	m := make(map[int]struct{}, 0)
	for _, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{v, target - v}
		} else {
			m[v] = struct{}{}
		}
	}
	return []int{}
}
