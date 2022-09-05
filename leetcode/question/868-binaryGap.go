package question

//868. 二进制间距
func binaryGap(n int) int {
	max := 0
	for i, j := 0, -1; n > 0; i++ {
		//如果等于1
		if n&1 == 1 {
			//减去，判断大小
			if j != -1 {
				if i-j > max {
					max = i - j
				}
			}
			j = i
		}
		n >>= 1
	}
	return max
}
