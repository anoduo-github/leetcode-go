package offer

//剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 == 0 {
			if nums[j]%2 == 1 {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
			j--
		} else {
			i++
		}
	}
	return nums
}
