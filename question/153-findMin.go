package question

//153. 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		pivot := left + (right-left)/2
		if nums[pivot] < nums[right] {
			right = pivot
		} else {
			left = pivot + 1
		}
	}
	return nums[left]
}
