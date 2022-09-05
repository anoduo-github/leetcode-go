package question

//209. 长度最小的子数组
func minSubArrayLen(target int, nums []int) int {
	index, last := 0, 0
	sum := nums[0]
	min := 0
	for last < len(nums) {
		if sum >= target {
			temp := last - index + 1
			if min > temp {
				min = temp
			} else if min == 0 {
				min = temp
			}
			sum -= nums[index]
			index++
		} else {
			last++
			if last < len(nums) {
				sum += nums[last]
			}
		}
	}
	return min
}
