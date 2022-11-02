package huawei

import (
	"fmt"
	"strings"
)

func HJ65(short, long string) {
	max := 0
	res := ""
	//遍历短的
	for i := 0; i < len(short)-max; i++ {
		for j := len(short) - 1; j > i; j-- {
			if strings.Contains(long, short[i:j+1]) {
				if max < j-i+1 {
					max = j - i + 1
					res = short[i : j+1]
					break
				}
			}
		}
	}
	fmt.Println(res)
}
