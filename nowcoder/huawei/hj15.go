package huawei

import "fmt"

//输入一个int型的正整数，计算出该int型数据在内存中存储时1的个数。

func hj15() {
	var num int
	fmt.Scanln(&num)
	//方法1.位运算， 每次数字减一，之后与原数字与，直到结果为0 原1100，减一，1011，则1100&1011=1000，重复，0111&1000=0000，2次
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	fmt.Println(count)
}
