package huawei

import (
	"fmt"
	"strconv"
	"strings"
)

//HJ90 验证ipv4地址
func HJ90() {
	var s string
	fmt.Scanln(&s)
	arr := strings.Split(s, ".")
	if len(arr) != 4 {
		fmt.Println("NO")
		return
	}
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		if !check(temp) {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

func check(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}
	if l == 1 {
		return s[0] >= '0' && s[0] <= '9'
	}
	if l == 2 {
		return (s[0] >= '1' && s[0] <= '9') && (s[1] >= '0' && s[1] <= '9')
	}
	if l == 3 {
		if (s[0] >= '1' && s[0] <= '2') && (s[1] >= '0' && s[1] <= '9') && (s[2] >= '0' && s[2] <= '9') {
			n, _ := strconv.Atoi(s)
			return n <= 255
		}
		return false
	}
	return false
}
