package question

//167. 两数之和 II - 输入有序数组
func twoSum167(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		temp := numbers[i] + numbers[j]
		if temp > target {
			//缩小右边
			j--
		} else if temp < target {
			//缩小左边
			i++
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return []int{0, 0}
}
