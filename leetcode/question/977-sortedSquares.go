package question

//977. 有序数组的平方
func sortedSquares(nums []int) []int {
	cur := len(nums) - 1
	res := make([]int, cur+1)
	index, last := 0, cur
	for index <= last && cur >= 0 {
		left := nums[index] * nums[index]
		right := nums[last] * nums[last]
		if left > right {
			res[cur] = left
			index++
		} else {
			res[cur] = right
			last--
		}
		cur--
	}
	return res
}
