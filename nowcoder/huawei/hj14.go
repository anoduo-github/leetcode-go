package huawei

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

//HJ41 称砝码
func HJ41() {
	//获取参数
	reader := bufio.NewScanner(os.Stdin)
	msgs := make([]string, 0)
	for reader.Scan() {
		s := reader.Text()
		msgs = append(msgs, s)
	}
	//将所有的砝码放入一个切片里
	n, _ := strconv.Atoi(msgs[0])
	arr1 := strings.Split(msgs[1], " ")
	arr2 := strings.Split(msgs[2], " ")
	temp := make([]int, 0)
	for i := 0; i < n; i++ {
		m, _ := strconv.Atoi(string(arr1[i]))
		x, _ := strconv.Atoi(string(arr2[i]))
		for j := 0; j < x; j++ {
			temp = append(temp, m)
		}
	}
	//统计重量
	set := make(map[int]struct{})
	set[0] = struct{}{}
	for i := 0; i < len(temp); i++ {
		backup := make([]int, 0)
		for i, _ := range set {
			backup = append(backup, i)
		}
		for _, v := range backup {
			set[v+temp[i]] = struct{}{}
		}
	}
	//遍历
	count := 1
	for range set {
		count++
	}
	fmt.Println(count)
}
