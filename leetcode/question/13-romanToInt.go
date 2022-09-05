package question

//13. 罗马数字转整数
func romanToInt(s string) int {
	//建立对照表
	romanMap := make(map[byte]int, 0)
	romanMap['I'] = 1
	romanMap['V'] = 5
	romanMap['X'] = 10
	romanMap['L'] = 50
	romanMap['C'] = 100
	romanMap['D'] = 500
	romanMap['M'] = 1000

	res := 0

	//从左到右遍历字符串（两两配对）
	for i := 0; i < len(s); i++ {
		leftInt := romanMap[s[i]] //左边的数字
		rightInt := 0             //右边的数字
		if i < len(s)-1 {
			rightInt = romanMap[s[i+1]]
		}

		//如果左边的数字大于右边,即为加
		if leftInt >= rightInt {
			res += leftInt
		} else {
			res -= leftInt
		}
	}
	return res
}
