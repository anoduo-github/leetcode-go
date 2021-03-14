package huawei

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* 将一个英文语句以单词为单位逆序排放。例如“I am a boy”，逆序排放后为“boy a am I”
所有单词之间用一个空格隔开，语句中除了英文字母外，不再包含其他字符 */

func hj13() {
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	msg = strings.TrimSpace(msg)
	arr := strings.Split(msg, " ")
	res := arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		res = res + " " + arr[i]
	}
	fmt.Println(res)
}
