package question

import sss "sort"

//976. 三角形的最大周长
func largestPerimeter(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	//先排序
	sss.Ints(nums)
	//从右开始
	last := len(nums) - 1
	for last >= 2 {
		i, j := last-2, last-1
		if nums[i]+nums[j] > nums[last] {
			//符合条件
			return nums[i] + nums[j] + nums[last]
		} else {
			last--
		}
	}
	return 0
}
