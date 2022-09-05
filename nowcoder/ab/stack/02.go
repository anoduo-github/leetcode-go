package stack

//AB2 栈的压入、弹出序列
func IsPopOrder(pushV []int, popV []int) bool {
	//辅助栈
	s := make([]int, 0)
	//拿辅助栈和输出序列对比
	a := 0    //输入下标
	b := 0    //输出下标
	size := 0 //辅助栈大小
	for a < len(pushV) {
		s = append(s, pushV[a])
		size++
		a++
		//判断此时的栈顶是否符合popV的序列顺序
		for size >= 1 && b < len(popV) && s[size-1] == popV[b] {
			s = s[:size-1]
			b++
			size--
		}
	}
	return b == len(popV)
}
