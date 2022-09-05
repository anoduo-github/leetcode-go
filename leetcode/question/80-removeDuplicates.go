package question

//80. 删除有序数组中的重复项 II
func removeDuplicates2(nums []int) int {
	//小于等于2不用考虑
	if len(nums) < 3 {
		return len(nums)
	}
	//双指针
	left, right := 2, 2
	for right < len(nums) {
		if nums[left-2] != nums[right] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
