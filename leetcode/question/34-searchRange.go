package question

//34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
		return []int{-1, -1}
	}
	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}
	//二分
	index, last := 0, len(nums)-1
	for index <= last {
		mid := (index + last) / 2
		if nums[mid] == target {
			left, right := mid-1, mid+1
			for left >= 0 {
				if nums[left] == target {
					left--
				} else {
					break
				}
			}
			for right < len(nums) {
				if nums[right] == target {
					right++
				} else {
					break
				}
			}
			return []int{left + 1, right - 1}
		} else if nums[mid] < target {
			index = mid + 1
		} else {
			last = mid - 1
		}
	}
	return []int{-1, -1}
}
