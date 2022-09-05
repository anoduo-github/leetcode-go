package question

//20. 有效的括号
func isValid(s string) bool {
	stack := make([]byte, 0)
	arr := []byte(s)
	for _, v := range arr {
		if len(stack) == 0 {
			stack = append(stack, v)
		} else {
			switch string(v) {
			case "(", "{", "[":
				stack = append(stack, v)
			case ")":
				if string(stack[len(stack)-1]) == "(" {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, v)
				}
			case "}":
				if string(stack[len(stack)-1]) == "{" {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, v)
				}
			case "]":
				if string(stack[len(stack)-1]) == "[" {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, v)
				}
			}
		}
	}
	return len(stack) == 0
}
