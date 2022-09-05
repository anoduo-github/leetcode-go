package question

//189. 轮转数组
func rotate(nums []int, k int) {
	res := make([]int, len(nums))
	for i, v := range nums {
		res[(i+k)%len(nums)] = v
	}
	copy(nums, res)
}
