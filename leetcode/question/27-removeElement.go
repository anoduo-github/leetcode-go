package question

//27. 移除元素
func removeElement(nums []int, val int) int {
	count := 0
	for _, v := range nums {
		if v != val {
			nums[count] = v
			count++
		}
	}
	return count
}
