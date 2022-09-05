package primary

//LC 合并两个有序数组
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
}

/* LC 第一个错误的版本， start */
func firstBadVersion(n int) int {
	x, y := 1, n
	mid := (x + y) / 2
	if isBadVersion(mid) {
		y = mid
	} else {
		x = mid + 1
	}
	return search(x, y)
}

func search(x, y int) int {
	if x == y {
		return x
	}
	mid := (x + y) / 2
	if isBadVersion(mid) {
		y = mid
	} else {
		x = mid + 1
	}
	return search(x, y)
}

/* LC 第一个错误的版本， end */

//提供的接口
func isBadVersion(n int) bool {
	/*
		逻辑代码
	*/
	// return false or
	return true
}
