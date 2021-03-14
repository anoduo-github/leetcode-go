package huawei

import (
	"bufio"
	"fmt"
	"os"
)

//字符串分隔
func hj4() {
	reader := bufio.NewReader(os.Stdin)
	chs := make([]string, 0)
	for {
		br, _, _ := reader.ReadLine()
		if len(br) == 0 {
			break
		}
		msg := string(br)
		if len(msg) <= 8 {
			count := 8 - len(msg)
			for count > 0 {
				msg = msg + "0"
				count--
			}
			chs = append(chs, msg)
		} else {
			a := len(msg) / 8
			b := len(msg) % 8
			for a > 0 {
				temp := msg[:8]
				chs = append(chs, temp)
				msg = msg[8:]
				a--
			}
			for b > 0 && 8-b > 0 {
				msg = msg + "0"
				b++
			}
			chs = append(chs, msg)
		}
	}
	for _, v := range chs {
		fmt.Println(v)
	}
}
