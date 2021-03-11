package practice

import (
	"fmt"
	"strconv"
)

/* 任给一个数组，元素有20M，1T，300G之类的，其中1T=1000G，1G=1000M
按从小到大输出结果
例如：输入：3
20M
1T
300G
输出：
20M
300G
1T */

func pr1() {
	//接受数量
	var num int
	fmt.Println("请输入数量:")
	fmt.Scanln(&num)
	if num == 0 {
		fmt.Println("无参数")
		return
	}
	//存放参数
	nums := make([]string, 0)
	fmt.Println("请输入参数:")
	for i := 0; i < num; i++ {
		var s string
		fmt.Scanln(&s)
		if len(s) == 0 || len(s) == 1 {
			fmt.Println("error: 参数缺失")
			return
		}
		switch string(s[len(s)-1]) {
		case "M", "G", "T":
			_, err := strconv.ParseInt(s[:len(s)-1], 0, 64)
			if err != nil {
				fmt.Println("参数格式错误")
				return
			}
			nums = append(nums, s)
		default:
			fmt.Println("error: 参数格式错误")
			return
		}
	}
	//排序
	sort(nums)
	fmt.Println("结果:")
	for _, v := range nums {
		fmt.Println(v)
	}
}

func sort(target []string) {
	for i := 0; i < len(target)-1; i++ {
		for j := 0; j < len(target)-i-1; j++ {
			if stoi(target[j]) > stoi(target[j+1]) {
				target[j], target[j+1] = target[j+1], target[j]
			}
		}
	}
}

func stoi(s string) int64 {
	size := int64(0)
	switch string(s[len(s)-1]) {
	case "M":
		size = 1
	case "G":
		size = 1000
	case "T":
		size = 1000000
	}
	n, _ := strconv.ParseInt(s[:len(s)-1], 0, 64)
	return size * n
}
