package question

//821. 字符的最短距离
func shortestToChar(s string, c byte) []int {
	length := len(s)
	res := make([]int, length)

	//初始化c的下标
	index := -length //这是表示c不存在的情况下
	//从左比较,存在左边部分值未正确计算出距离，因为初始c的下标
	for i := 0; i < length; i++ {
		if s[i] == c {
			index = i
		}
		res[i] = i - index
	}

	//从右比较，右边部分不正确，根据上面左边遍历的左右互补，即为结果
	index = 2 * length
	for j := length - 1; j >= 0; j-- {
		if s[j] == c {
			index = j
		}
		if res[j] > index-j {
			res[j] = index - j
		}
	}

	return res
}
