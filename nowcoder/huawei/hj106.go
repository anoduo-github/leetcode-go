package huawei

import (
	"bufio"
	"fmt"
	"os"
)

//将一个字符串str的内容颠倒过来，并输出。str的长度不超过100个字符。

func hj106() {
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	if len(msg) == 0 {
		fmt.Println(msg)
	}
	arr := []rune(msg)
	left, right := 0, len(arr)-1
	for left <= right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	fmt.Println(string(arr))
}
