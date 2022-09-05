package question

//946. 验证栈序列
func validateStackSequences(pushed []int, popped []int) bool {
	//临时栈
	temp := make([]int, 0)
	//pushed下标，popped下标
	i, j := 0, 0
	for ; i < len(pushed); i++ {
		temp = append(temp, pushed[i])
		//判断临时栈的栈顶元素和popped的第j个元素是否相同，同就pop出去
		for len(temp) > 0 && j < len(popped) {
			if popped[j] == temp[len(temp)-1] {
				if len(temp) == 1 {
					temp = make([]int, 0)
				} else {
					temp = temp[0 : len(temp)-1]
				}
				//同时j后移一位
				j++
			} else {
				break
			}
		}
	}
	return len(temp) == 0
}
