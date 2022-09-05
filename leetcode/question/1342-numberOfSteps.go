package question

//1342. 将数字变成 0 的操作次数
func numberOfSteps(num int) int {
	count := 0
	for num > 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}
		count++
	}
	return count
}
