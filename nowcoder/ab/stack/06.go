package stack

import "strconv"

//错误

//AB6 表达式求值
func solve(s string) int {
	//设计一个map存储符号和优先级
	m := make(map[byte]int, 0)
	m['+'] = 0
	m['-'] = 0
	m['*'] = 1
	m['('] = -1
	m[')'] = 2
	//建立一个数栈和一个符号栈
	res := make([]int, 0)
	r := 0
	symbol := make([]byte, 0)
	ss := 0
	//遍历
	temp := 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			if temp < i {
				//表示之前的都是数
				n, _ := strconv.Atoi(s[temp:i])
				res = append(res, n)
				r++
			}
			temp = i + 1
			//判断符号优先级
			if ss == 0 {
				symbol = append(symbol, s[i])
				ss++
			} else if m[s[i]]-m[symbol[ss-1]] >= 0 {
				if s[i] == ')' {
					//弹出符号，,取数进行计算,直到遇到（
					t := ss - 1
					for t >= 0 && symbol[t] != '(' {
						var num int
						switch symbol[t] {
						case '+':
							num = res[r-2] + res[r-1]
						case '-':
							num = res[r-2] - res[r-1]
						case '*':
							num = res[r-2] * res[r-1]
						}
						res[r-2] = num
						res = res[:r-1]
						r--
						t--
					}
					if t <= 0 {
						symbol = make([]byte, 0)
						ss = 0
					} else {
						symbol = symbol[:t]
						ss = t
					}
				} else {
					symbol = append(symbol, s[i])
					ss++
				}
			} else {
				if symbol[ss-1] != '(' && s[i] != '(' {
					//使用
					var num int
					switch s[i] {
					case '+':
						num = res[r-2] + res[r-1]
					case '-':
						num = res[r-2] - res[r-1]
					case '*':
						num = res[r-2] * res[r-1]
					}
					res[r-2] = num
					res = res[:r-1]
					r--
				} else {
					symbol = append(symbol, s[i])
					ss++
				}
			}
		}
	}
	if temp < len(s) {
		n, _ := strconv.Atoi(s[temp:])
		res = append(res, n)
		r++
	}
	//将符号栈剩余的拿出进行运算
	for i := ss - 1; i >= 0; i-- {
		var num int
		switch symbol[i] {
		case '+':
			num = res[r-2] + res[r-1]
		case '-':
			num = res[r-2] - res[r-1]
		case '*':
			num = res[r-2] * res[r-1]
		}
		res[r-2] = num
		res = res[:r-1]
		r--
	}

	return res[r-1]
}
