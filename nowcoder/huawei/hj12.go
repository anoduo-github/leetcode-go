package huawei

import "fmt"

//接受一个只包含小写字母的字符串，然后输出该字符串反转后的字符串。（字符串长度不超过1000）

func hj12() {
	var str string
	fmt.Scanln(&str)
	arr := []byte(str)
	left := 0
	right := len(arr) - 1
	for left <= right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	fmt.Println(string(arr))
}
