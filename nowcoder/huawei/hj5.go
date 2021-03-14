package huawei

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//写出一个程序，接受一个十六进制的数，输出该数值的十进制表示。
func hj5() {
	reader := bufio.NewReader(os.Stdin)
	arr := make([]string, 0)
	for {
		rl, _, _ := reader.ReadLine()
		msg := string(rl)
		if len(msg) == 0 {
			break
		}
		arr = append(arr, msg)
	}
	for _, v := range arr {
		v = v[2:]
		//总和
		sum := 0
		//幂
		counts := 0
		//数字
		num := 0
		for i := len(v) - 1; i >= 0; i-- {
			switch string(v[i]) {
			case "A":
				num = 10
			case "B":
				num = 11
			case "C":
				num = 12
			case "D":
				num = 13
			case "E":
				num = 14
			case "F":
				num = 15
			default:
				n, _ := strconv.Atoi(string(v[i]))
				num = n
			}
			temp := 16
			//求幂
			count := counts
			if count == 0 {
				temp = 1
			} else {
				for count > 1 {
					temp = temp * 16
					count--
				}
			}
			sum = sum + num*temp
			counts++
		}
		fmt.Println(sum)
	}
}
