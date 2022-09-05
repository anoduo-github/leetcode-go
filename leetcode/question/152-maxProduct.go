package question

//152. 乘积最大子数组
func maxProduct152(nums []int) int {
	max := nums[0]
	res := 1
	//从左往右
	for _, v := range nums {
		res *= v
		if res > max {
			max = res
		}
		if v == 0 {
			res = 1
		}
	}
	res = 1
	//从右往左
	for i := len(nums) - 1; i >= 0; i-- {
		res *= nums[i]
		if res > max {
			max = res
		}
		if nums[i] == 0 {
			res = 1
		}
	}
	return max
}
