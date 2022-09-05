package question

//942. 增减字符串匹配

func diStringMatch(s string) []int {
	arr := []rune(s)
	max := len(arr)
	//存整数集合
	nums := make([]int, max+1)
	for i := 0; i <= max; i++ {
		nums[i] = i
	}
	res := make([]int, max+1)
	//遍历
	for i := 0; i < max; i++ {
		if arr[i] == 'I' {
			//赋值最小的
			res[i] = nums[0]
			//缩小范围
			nums = nums[1:]
		} else {
			//赋值最大的
			res[i] = nums[len(nums)-1]
			//缩小范围
			nums = nums[0 : len(nums)-1]
		}
	}
	res[max] = nums[0]
	return res
}
