package question

//350. 两个数组的交集 II
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, 0)
	if len(nums1) < len(nums2) {
		return intersect(nums2, nums1)
	}
	for _, v := range nums1 {
		m[v]++
	}
	res := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if _, ok := m[nums2[i]]; ok {
			if m[nums2[i]] > 0 {
				res = append(res, nums2[i])
				m[nums2[i]]--
			}
		}
	}
	return res
}
