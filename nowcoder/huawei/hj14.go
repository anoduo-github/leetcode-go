package huawei

import (
	"fmt"
)

//给定n个字符串，请对n个字符串按照字典序排列。

func hj14() {
	var n int
	fmt.Scanln(&n)
	res := make([]string, 0)
	for n > 0 {
		var str string
		fmt.Scanln(&str)
		res = append(res, str)
		n--
	}
	for i := 0; i < len(res)-1; i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i] > res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}
}
