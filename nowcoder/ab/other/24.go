package other

//AB24 二分查找-I
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	//双指针
	left := 0
	right := len(nums) - 1
	mid := len(nums) / 2
	//判断
	for left < right {
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
			mid = (left + right) / 2
		} else {
			right = mid - 1
			mid = (left + right) / 2
		}
	}
	return -1
}
