package question

//1389. 按既定顺序创建目标数组
func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(nums))
	for k, v := range index {
		//这一步可以理解为方便插入，其余数据不变
		copy(res[v+1:], res[v:]) //=> 将切片[v+1:]复制到[v:]中，即从v+1开始的值均不变，方便v处插入值
		res[v] = nums[k]
	}
	return res
}
