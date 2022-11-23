package huawei

import "fmt"

//HJ21 简单密码
func HJ21() {
	var s string
	fmt.Scanln(&s)
	arr := []byte(s)
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 'a' && arr[i] <= 'z' {
			switch arr[i] {
			case 'a', 'b', 'c':
				arr[i] = '2'
			case 'd', 'e', 'f':
				arr[i] = '3'
			case 'g', 'h', 'i':
				arr[i] = '4'
			case 'j', 'k', 'l':
				arr[i] = '5'
			case 'm', 'n', 'o':
				arr[i] = '6'
			case 'p', 'q', 'r', 's':
				arr[i] = '7'
			case 't', 'u', 'v':
				arr[i] = '8'
			case 'w', 'x', 'y', 'z':
				arr[i] = '9'
			}
		} else if arr[i] >= 'A' && arr[i] <= 'Z' {
			if arr[i] == 'Z' {
				arr[i] = 'a'
			} else {
				arr[i] = arr[i] + 'a' - 'A' + 1
			}
		}
	}
	fmt.Println(string(arr))
}
