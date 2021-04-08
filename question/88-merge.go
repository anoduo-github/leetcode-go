package question

//88. 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	//设置数组下标，从后往前遍历
	right1, right2, right := m-1, n-1, m+n-1
	for right2 >= 0 {
		if right1 >= 0 && nums1[right1] > nums2[right2] {
			nums1[right] = nums1[right1]
			right1--
		} else {
			nums1[right] = nums2[right2]
			right2--
		}
		right--
	}
	/* //将nums2放在nums1末尾
	i := m
	j := 0
	for i < len(nums1) && j < len(nums2) {
		nums1[i] = nums2[j]
		i++
		j++
	}
	//排序
	for k := 0; k < len(nums1)-1; k++ {
		for m := k + 1; m < len(nums1); m++ {
			if nums1[k] > nums1[m] {
				nums1[k], nums1[m] = nums1[m], nums1[k]
			}
		}
	} */
}
