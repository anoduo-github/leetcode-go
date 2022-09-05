package question

//162. 寻找峰值
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	index, last := 0, len(nums)-1
	for index < last {
		mid := (index + last) / 2
		if nums[mid] > nums[mid+1] {
			last = mid
		} else {
			index = mid + 1
		}
	}
	return index
}
