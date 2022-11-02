package question

import "sort"

//1636. 按照频率将数组升序排序
func frequencySort(nums []int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		return m[a] < m[b] || m[a] == m[b] && a > b
	})
	return nums
}
