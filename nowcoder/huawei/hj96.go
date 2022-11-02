package huawei

import "fmt"

//HJ96 表示数字
func HJ96(s string) {
	arr := []byte(s)
	res := make([]byte, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] >= '0' && arr[i] <= '9' {
			if i == 0 || (i > 0 && (arr[i-1] < '0' || arr[i-1] > '9')) {
				res = append(res, '*')
			}
			res = append(res, arr[i])
			if (i < len(arr)-1 && (arr[i+1] < '0' || arr[i+1] > '9')) || i == len(arr)-1 {
				res = append(res, '*')
			}
		} else {
			res = append(res, arr[i])
		}
	}
	fmt.Println(string(res))
}
