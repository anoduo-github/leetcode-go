package question

import "strconv"

//150. 逆波兰表达式求值
func evalRPN(tokens []string) int {
	//准备一个数栈
	nums := make([]int, 0)
	//从左至右遍历,将数存入数栈
	for _, v := range tokens {
		switch v {
		//遇到运算符，取出两个数进行运算，第二个数放在运算符的左边
		case "+":
			temp := nums[len(nums)-2] + nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			nums = append(nums, temp)
		case "-":
			temp := nums[len(nums)-2] - nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			nums = append(nums, temp)
		case "*":
			temp := nums[len(nums)-2] * nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			nums = append(nums, temp)
		case "/":
			temp := nums[len(nums)-2] / nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			nums = append(nums, temp)
		default:
			//遇到数就放入
			temp, _ := strconv.Atoi(v)
			nums = append(nums, temp)
		}
	}
	return nums[len(nums)-1]
}
