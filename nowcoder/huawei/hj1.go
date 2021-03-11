package huawei

import (
	"bufio"
	"os"
	"strings"
)

//getLastLen 计算字符串最后一个单词的长度
func getLastLen() int {
	//获取输入
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	msg = strings.TrimSpace(msg)
	arr := strings.Split(msg, " ")
	if len(arr) == 0 {
		return 0
	}
	return len(arr[len(arr)-1])
}
