package question

//1422. 分割字符串的最大得分
func maxScore(s string) int {
	num0 := 0
	num1 := 0
	for _, v := range s {
		if v == '0' {
			num0++
		}
		if v == '1' {
			num1++
		}
	}
	//遍历
	index := 0
	temp0 := 0
	temp1 := 0
	max := 0
	for index < len(s)-1 {
		if s[index] == '0' {
			temp0++
		}
		if s[index] == '1' {
			temp1++
		}
		temp := temp0 + num1 - temp1
		if max < temp {
			max = temp
		}
		index++
	}
	return max
}
