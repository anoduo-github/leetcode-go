package huawei

import "fmt"

//HJ29 字符串加解密
func HJ29() {
	/* var x string
	fmt.Scanln(&x)
	var y string
	fmt.Scanln(&y) */
	x := "2OA92AptLq5G1lW8564qC4nKMjv8C"
	y := "B5WWIj56vu72GzRja7j5"

	fmt.Println(en(x))
	fmt.Println(dec(y))
}

//加密
func en(s string) string {
	arr := []byte(s)
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		if temp >= 'a' && temp <= 'z' {
			if temp == 'z' {
				arr[i] = 'A'
			} else {
				arr[i] = temp - 31
			}
		} else if temp >= 'A' && temp <= 'Z' {
			if temp == 'Z' {
				arr[i] = 'a'
			} else {
				arr[i] = temp + 33
			}
		} else {
			if temp == '9' {
				arr[i] = '0'
			} else {
				arr[i] = temp + 1
			}
		}
	}
	return string(arr)
}

//解密
func dec(s string) string {
	arr := []byte(s)
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		if temp >= 'a' && temp <= 'z' {
			if temp == 'a' {
				arr[i] = 'Z'
			} else {
				arr[i] = temp - 33
			}
		} else if temp >= 'A' && temp <= 'Z' {
			if temp == 'A' {
				arr[i] = 'z'
			} else {
				arr[i] = temp + 31
			}
		} else {
			if temp == '0' {
				arr[i] = '9'
			} else {
				arr[i] = temp - 1
			}
		}
	}
	return string(arr)
}
