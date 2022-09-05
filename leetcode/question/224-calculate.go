package question

import "fmt"

//224. 基本计算器
func calculate(s string) int {
	res := 0
	//栈
	stack := make([]int, 0)
	//标志位
	tag := 1
	stack = append(stack, tag)
	//坐标
	i := 0
	for i < len(s) {
		switch string(s[i]) {
		case "+":
			tag = stack[len(stack)-1]
			i++
		case "-":
			tag = -stack[len(stack)-1]
			i++
		case "(":
			stack = append(stack, tag)
			i++
		case ")":
			stack = stack[:len(stack)-1]
			i++
		case " ":
			i++
		default:
			n := 0
			for i < len(s) {
				if s[i] >= '0' && s[i] <= '9' {
					n = n*10 + int(s[i]-'0')
					fmt.Println(n)
					i++
				} else {
					break
				}
			}
			res = res + tag*n
		}
	}
	return res
}
