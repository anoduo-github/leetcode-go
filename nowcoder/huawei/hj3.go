package huawei

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//hj3 去重和排序
func hj3() {
	reader := bufio.NewReader(os.Stdin)
	for {
		tag, _, _ := reader.ReadLine()
		if len(tag) == 0 {
			break
		}
		count, _ := strconv.Atoi(string(tag))
		mark := [1000]bool{}
		for i := 0; i < count; i++ {
			ns, _, _ := reader.ReadLine()
			n, _ := strconv.Atoi(string(ns))
			mark[n] = true
		}
		for k, v := range mark {
			if v {
				fmt.Println(k)
			}
		}
	}
}
