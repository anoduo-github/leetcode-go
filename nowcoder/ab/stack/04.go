package stack

import "strconv"

//AB4 逆波兰表达式求值
func evalRPN(tokens []string) int {
	//辅助栈
	//从左往右遍历，把数存入栈，遇到符号就取出第一个和第二个进行计算，如何再存入栈中
	stack := make([]int, 0)
	size := 0
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			res := stack[size-2] + stack[size-1]
			stack = stack[:size-2]
			stack = append(stack, res)
			size--
		case "-":
			res := stack[size-2] - stack[size-1]
			stack = stack[:size-2]
			stack = append(stack, res)
			size--
		case "*":
			res := stack[size-2] * stack[size-1]
			stack = stack[:size-2]
			stack = append(stack, res)
			size--
		case "/":
			res := stack[size-2] / stack[size-1]
			stack = stack[:size-2]
			stack = append(stack, res)
			size--
		default:
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
			size++
		}
	}
	return stack[size-1]
}
