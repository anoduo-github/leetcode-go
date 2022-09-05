package stack

// eliminate 消除重复的字符
func eliminate(s string) string {
	arr := []byte(s)
	//辅助栈
	stack := make([]byte, 0)
	size := 0
	//遍历
	for i := 0; i < len(arr); i++ {
		if size > 0 && stack[size-1] == arr[i] {
			//表示这两个相挨且相同
			//去除重复的，由于这个arr[i]未存入栈，所以只要把stack[size-1]弹出就行
			stack = stack[:size-1]
			size--
		} else {
			stack = append(stack, arr[i])
			size++
		}
	}
	if size == 0 {
		return "0"
	}
	return string(stack)
}
