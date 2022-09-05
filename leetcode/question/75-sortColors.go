package question

//75. 颜色分类
func sortColors(nums []int) {
	index, last := 0, len(nums)-1 //index指向0，last指向2
	i := 0
	for i <= last {
		if nums[i] == 0 {
			nums[i], nums[index] = nums[index], nums[i]
			index++
			i++
			continue
		}
		if nums[i] == 1 {
			i++
			continue
		}
		if nums[i] == 2 {
			nums[i], nums[last] = nums[last], nums[i]
			last--
		}
	}
}
