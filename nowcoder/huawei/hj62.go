package huawei

import "fmt"

//输入一个正整数，计算它在二进制下的1的个数。

func hj62() {
	nums := make([]int, 0)
	for {
		var num int
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		nums = append(nums, num)
	}
	for _, v := range nums {
		count := 0
		for v > 0 {
			v = v & (v - 1)
			count++
		}
		fmt.Println(count)
	}
}
