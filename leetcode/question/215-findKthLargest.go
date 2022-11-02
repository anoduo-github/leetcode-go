package question

import (
	"sort"
)

//215. 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[k-1]
}
