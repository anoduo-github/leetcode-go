package question

//283. 移动零
func moveZeroes(nums []int) {
	//遍历数组，将不为0的往前面放，剩下的就是为0的
	index := 0
	for _, v := range nums {
		if v != 0 {
			nums[index] = v
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}
