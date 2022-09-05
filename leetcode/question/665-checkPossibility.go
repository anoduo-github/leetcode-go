package question

//665. 非递减数列
func checkPossibility(nums []int) bool {
	//思路：3个一组进行判断，保证每组都是递增
	//数组长度
	l := len(nums)
	if l <= 2 {
		return true
	}
	//初始下标，表示每组中间的那个值
	i := 1
	//计数器,初始为1，表示只1次互换的机会
	count := 1
	for i <= l-2 {
		//当左边大于中间
		if nums[i-1] > nums[i] {
			//当中间大于右边，超过了1次互换，错误
			if nums[i] > nums[i+1] {
				return false
			}
			//当中间小于右边，让左边值等于中间值
			nums[i-1] = nums[i]
			count--
		} else {
			//当左边小于中间且左边大于中间时
			if nums[i] > nums[i+1] {
				//当左边大于右边
				if nums[i-1] > nums[i+1] {
					//右边值等于中间值
					nums[i+1] = nums[i]
				} else {
					//中间值等于左边值
					nums[i] = nums[i-1]
				}
				count--
			}
		}
		//count<0表示交换次数超过了一次
		if count < 0 {
			return false
		}
		i++
	}
	return true
}
