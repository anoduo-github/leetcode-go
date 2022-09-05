package question

//806. 写字符串需要的行数
func numberOfLines(widths []int, s string) []int {
	amount := 0 //每行占用的数量
	lines := 1  //行数,因为题干给的是字符串长度至少为1，至少占一行
	for _, v := range s {
		//计算出对应的widths数组下标，因为widths数组一共26个，一一对应字母
		index := v - 'a'
		//累加
		amount += widths[index]
		if amount > 100 {
			amount = widths[index]
			lines++
		}
	}
	return []int{lines, amount}
}
