package stack

import (
	"strconv"
)

func solves(s string) int {
	res := infixToSuffix(s)
	temp := make([]int, 0)
	size := 0
	for i := 0; i < len(res); i++ {
		switch res[i] {
		case "+":
			num := temp[size-2] + temp[size-1]
			temp = temp[:size-1]
			temp[size-2] = num
			size--
		case "-":
			num := temp[size-2] - temp[size-1]
			temp = temp[:size-1]
			temp[size-2] = num
			size--
		case "*":
			num := temp[size-2] * temp[size-1]
			temp = temp[:size-1]
			temp[size-2] = num
			size--
		default:
			num, _ := strconv.Atoi(res[i])
			temp = append(temp, num)
			size++
		}
	}
	return temp[size-1]
}

func infixToSuffix(s string) []string {
	//设计一个map存储符号和优先级
	m := make(map[byte]int, 0)
	m['+'] = 0
	m['-'] = 0
	m['*'] = 1
	m['('] = 2
	m[')'] = 2
	//建立一个结果栈和一个符号栈
	res := make([]string, 0)
	symbol := make([]byte, 0)
	size := 0
	//遍历
	temp := 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			if temp < i {
				//表示之前的都是数
				res = append(res, s[temp:i])
			}
			temp = i + 1
			//判断符号优先级
			if size == 0 {
				symbol = append(symbol, s[i])
				size++
			} else if m[s[i]]-m[symbol[size-1]] >= 0 {
				if s[i] == ')' {
					//弹出符号，直到遇到（
					t := size - 1
					for t >= 0 && symbol[t] != '(' {
						res = append(res, string(symbol[t]))
						t--
					}
					if t <= 0 {
						symbol = make([]byte, 0)
						size = 0
					} else {
						symbol = symbol[:t]
						size = t
					}
				} else {
					symbol = append(symbol, s[i])
					size++
				}
			} else {
				if symbol[size-1] != '(' {
					//将栈顶符号放入结果中
					res = append(res, string(symbol[size-1]))
					//再放入符号栈
					symbol[size-1] = s[i]
				} else {
					symbol = append(symbol, s[i])
					size++
				}
			}
		}
	}
	if temp < len(s) {
		res = append(res, s[temp:])
	}
	//将符号栈剩余的放入结果栈
	for i := size - 1; i >= 0; i-- {
		res = append(res, string(symbol[i]))
	}

	return res
}
