package huawei

import (
	"fmt"
	"strings"
)

//HJ66 配置文件恢复
func HJ66(s string, m map[string]string) {
	arr := strings.Split(s, " ")
	if len(arr) > 2 {
		fmt.Println("unknown command")
		return
	}
	res := make([]string, 0)
	for k, _ := range m {
		temp := strings.Split(k, " ")
		if len(arr) != len(temp) {
			continue
		}
		count := 0
		for i := 0; i < len(arr); i++ {
			a := arr[i]
			b := temp[i]
			if len(a) <= len(b) && a == b[0:len(a)] {
				count++
			}
		}
		if count == len(arr) {
			res = append(res, k)
		}
	}
	if len(res) == 0 || len(res) > 1 {
		fmt.Println("unknown command")
	} else {
		fmt.Println(m[res[0]])
	}
}
