package huawei

import (
	"fmt"
	"strconv"
)

func Calculator(s string) {
	nums := make([]int, 0)
	symbol := make([]byte, 0)

	for i, j := 0, 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if i == len(s)-1 {
				temp, _ := strconv.Atoi(s[j : i+1])
				nums = append(nums, temp)
				break
			}
			continue
		} else {
			if i > j {
				temp, _ := strconv.Atoi(s[j:i])
				nums = append(nums, temp)
			}
			j = i + 1
			b := s[i]
			switch b {
			case '(':
				symbol = append(symbol, b)
			case ')':
				//取出符号进行计算
				for i := len(symbol) - 1; i >= 0; i-- {
					if symbol[i] == '(' {
						symbol = symbol[0:i]
						break
					}
					var res int
					if len(nums) == 1 {
						res = caculate(0, nums[0], symbol[i])
						nums = make([]int, 0)
					} else {
						res = caculate(nums[len(nums)-2], nums[len(nums)-1], symbol[i])
						nums = nums[0 : len(nums)-2]
					}
					nums = append(nums, res)
				}
			case '+', '-':
				if len(symbol) == 0 {
					symbol = append(symbol, b)
				} else {
					i := len(symbol) - 1
					//比较等级
					for i >= 0 {
						if symbol[i] == '(' || symbol[i] == '+' || symbol[i] == '-' {
							symbol = symbol[0 : i+1]
							symbol = append(symbol, b)
							break
						}
						if symbol[i] == '*' || symbol[i] == '/' {
							//计算
							res := caculate(nums[len(nums)-2], nums[len(nums)-1], symbol[i])
							nums = nums[0 : len(nums)-2]
							nums = append(nums, res)
						}
						i--
					}
					if i < 0 {
						symbol = make([]byte, 0)
						symbol = append(symbol, b)
					}
				}
			case '*', '/':
				symbol = append(symbol, b)
			}
		}
	}

	//计算
	for i := len(symbol) - 1; i >= 0; i-- {
		var res int
		if len(nums) == 1 {
			res = caculate(0, nums[0], symbol[i])
			nums = make([]int, 0)
		} else {
			res = caculate(nums[len(nums)-2], nums[len(nums)-1], symbol[i])
			nums = nums[0 : len(nums)-2]
		}
		nums = append(nums, res)
	}
	fmt.Println(nums[len(nums)-1])
}

func caculate(a, b int, s byte) int {
	switch s {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return 0
}
