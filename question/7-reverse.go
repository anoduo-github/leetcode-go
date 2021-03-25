package question

//7.整数反转
func reverse(x int) int {
	var res int
	for {
		if x != 0 {
			res = res*10 + x%10
			x = x / 10
		} else {
			break
		}
	}
	//res > 1<<31-1 || res < -1<<31
	if res > 2147483647 || res < -2147483648 {
		return 0
	}
	return res
}
