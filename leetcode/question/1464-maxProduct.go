package question

//1464. 数组中两元素的最大乘积
func maxProduct(nums []int) int {
	max := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			temp := (nums[i] - 1) * (nums[j] - 1)
			if temp > max {
				max = temp
			}
		}
	}
	return max
}
