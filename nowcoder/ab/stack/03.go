package stack

//AB3 有效括号序列
func isValid(s string) bool {
	//辅助栈
	stack := make([]rune, 0)
	size := 0
	//遍历
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		//判断压入的是否和栈顶的为一对
		if size > 0 {
			switch stack[size-1] {
			case '(':
				if arr[i] == ')' {
					stack = stack[:size-1]
					size--
				} else {
					stack = append(stack, arr[i])
					size++
				}
			case '[':
				if arr[i] == ']' {
					stack = stack[:size-1]
					size--
				} else {
					stack = append(stack, arr[i])
					size++
				}
			case '{':
				if arr[i] == '}' {
					stack = stack[:size-1]
					size--
				} else {
					stack = append(stack, arr[i])
					size++
				}
			}
		} else {
			stack = append(stack, arr[i])
			size++
		}
	}
	return size == 0
}
