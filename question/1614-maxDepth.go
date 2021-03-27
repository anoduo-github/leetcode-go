package question

//1614. 括号的最大嵌套深度
func maxDepth(s string) int {
	//从左至右遍历，遇到（加一，遇到）减一，记录最大数
	max := 0
	temp := 0
	for _, v := range s {
		switch string(v) {
		case "(":
			temp++
		case ")":
			temp--
		}
		if max < temp {
			max = temp
		}
	}
	return max
}
