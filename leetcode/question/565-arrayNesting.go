package question

//565. 数组嵌套
func arrayNesting(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			continue
		}
		cur := i
		count := 0
		for nums[cur] >= 0 {
			count++
			cur, nums[cur] = nums[cur], -1
		}
		if count > max {
			max = count
		}
	}
	return max
}
