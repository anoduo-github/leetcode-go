package huawei

import (
	"fmt"
	"strconv"
)

/* 输入一个整数，将这个整数以字符串的形式逆序输出
程序不考虑负数的情况，若数字含有0，则逆序形式也含有0，如输入为100，则输出为001 */

func hj11() {
	var num int
	fmt.Scanln(&num)
	//循环取模法
	/* res := ""
	for num > 0 {
		temp := num % 10
		s := strconv.Itoa(temp)
		res = res + s
		num = num / 10
	} */
	//双指针法
	arr := []byte(strconv.Itoa(num))
	left := 0
	right := len(arr) - 1
	for left <= right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	fmt.Println(string(arr))
}
