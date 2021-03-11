package question

import (
	"fmt"
	"strconv"
)

//227. 基本计算器 II
func Calculate2(s string) int {
	//转为逆波兰表达式
	res := revpo(s)
	fmt.Println(res)
	//执行逆波兰表达式
	sum := doRevpo(res)
	return sum
}

//doRevpo 执行逆波兰表达式
func doRevpo(s []string) int {
	//定义一个栈
	nums := make([]int, 0)
	//遍历
	for _, v := range s {
		switch v {
		case "+", "-", "*", "/":
			//从nums取两个数
			num1 := nums[len(nums)-1]
			num2 := nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			if v == "+" {
				nums = append(nums, num2+num1)
			} else if v == "-" {
				nums = append(nums, num2-num1)
			} else if v == "*" {
				nums = append(nums, num2*num1)
			} else {
				nums = append(nums, num2/num1)
			}
		default:
			//将字符串转化为int
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}
	}
	return nums[len(nums)-1]
}

//revpo 求逆波兰表达式
func revpo(s string) []string {
	//1.定义两个栈
	res := make([]string, 0)
	chs := make([]string, 0)
	//2.遍历字符串
	for i := 0; i < len(s); i++ {
		temp := string(s[i])
		switch temp {
		case "+", "-", "*", "/":
			//假如符号栈为空，则放入
			if len(chs) == 0 {
				chs = append(chs, temp)
				break
			}
			//假如符号栈栈顶元素为（，则直接放入
			if string(chs[len(chs)-1]) == "(" {
				chs = append(chs, temp)
				break
			}
			//小于则弹出栈顶元素到res中，循环执行，直到大于为止再放入
			for {
				if len(chs) == 0 || compare(temp) > compare(chs[len(chs)-1]) {
					//大于或者为空则直接放入
					chs = append(chs, temp)
					break
				}
				//小于
				res = append(res, chs[len(chs)-1])
				chs = chs[:len(chs)-1]
			}
		case "(":
			chs = append(chs, temp)
		case ")":
			//依次将符号栈中的元素弹出加入res，直到遇到（
			for {
				ch := chs[len(chs)-1]
				if len(chs) > 0 {
					chs = chs[:len(chs)-1]
				}
				if ch == "(" {
					break
				}
				res = append(res, ch)
			}
		case " ":
			continue
		default:
			//循环获取数字
			n := ""
			for {
				if s[i] >= '0' && s[i] <= '9' {
					n = n + string(s[i])
					i++
					//防止超过长度
					if i >= len(s) {
						i--
						break
					}
				} else {
					i--
					break
				}
			}
			res = append(res, n)
		}
	}
	//扫描完毕，将剩余的符号放入res中
	for j := len(chs) - 1; j >= 0; j-- {
		res = append(res, chs[j])
	}
	return res
}

//compare 比较优先性
func compare(s string) int {
	switch s {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}
