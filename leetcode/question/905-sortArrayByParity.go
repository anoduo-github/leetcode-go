package question

//905. 按奇偶排序数组
func sortArrayByParity(nums []int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left]%2 == 0 {
			left++
		} else {
			if nums[right]%2 == 0 {
				nums[left], nums[right] = nums[right], nums[left]
				left++
				right--
			} else {
				right--
			}
		}
	}
	return nums
}
