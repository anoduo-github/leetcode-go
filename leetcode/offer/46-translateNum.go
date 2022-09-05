package offer

import "strconv"

//剑指 Offer 46. 把数字翻译成字符串
func translateNum(num int) int {
	str := strconv.Itoa(num)
	one, two, res := 0, 0, 1
	for i := 0; i < len(str); i++ {
		one, two, res = two, res, 0
		res += two
		if i == 0 {
			continue
		}
		temp := str[i-1 : i+1]
		if temp >= "10" && temp <= "25" {
			res += one
		}
	}
	return res
}
