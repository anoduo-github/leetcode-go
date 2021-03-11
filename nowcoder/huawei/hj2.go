package huawei

import (
	"bufio"
	"os"
	"strings"
)

//getCharCount 获取某字母出现的次数
func getCharCount() int {
	reader := bufio.NewReader(os.Stdin)
	//接受第一个输入
	msg, _ := reader.ReadString('\n')
	msg = strings.ToUpper(msg)
	//接受第二个参数
	ch, _ := reader.ReadString('\n')
	ch = strings.TrimSpace(ch)
	if len(ch) == 0 {
		return 0
	}
	ch = strings.ToUpper(ch)
	//遍历比较
	count := 0
	for _, v := range msg {
		if string(v) == ch {
			count++
		}
	}
	return count
}
