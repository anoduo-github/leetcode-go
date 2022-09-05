package offer

//剑指 Offer 53 - I. 在排序数组中查找数字 I
func search(nums []int, target int) int {
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}
	return m[target]
}
